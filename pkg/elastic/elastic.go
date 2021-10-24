package elastic

import (
    "context"
    "fmt"
    "log"
    "os"
    "github.com/olivere/elastic/v7"
	"go-server-template/config"
)

var client *elastic.Client
var host = "http://"+projectConfig.AppConfig.ElasticConfig.HOST+":"+projectConfig.AppConfig.ElasticConfig.PORT+"/"

type ElesticInfo struct {
    Index string   `json:"index"`
    Type  string   `json:"type"`
    Id    string   `json:"id"`
}

type QueryCondition struct {
    Key string   
    Value  string   
}


//初始化
func InitElestic() {
    errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
    var err error
    client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
    if err != nil {
        panic(err)
    }
    info, code, err := client.Ping(host).Do(context.Background())
    if err != nil {
        panic(err)
    }
    fmt.Printf("Elasticsearch 返回的code %d 和版本 %s\n", code, info.Version.Number)

    esversion, err := client.ElasticsearchVersion(host)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Elasticsearch 版本 %s\n", esversion)

}

/*下面是简单的CURD*/

//创建
func Create(params *ElesticInfo, data interface{}) {
	put, err := client.Index().
		Index(params.Index).
		Type(params.Type).
		Id(params.Id).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put.Id, put.Index, put.Type)
}

//删除
func Delete(params *ElesticInfo) {
    res, err := client.Delete().Index(params.Index).
        Type(params.Type).
        Id(params.Id).
        Do(context.Background())
    if err != nil {
        println(err.Error())
        return
    }
    fmt.Printf("delete result %s\n", res.Result)
}

//修改
func Update(params *ElesticInfo, data map[string]interface{}) {
    res, err := client.Update().
        Index(params.Index).
        Type(params.Type).
        Id(params.Id).
        Doc(data).
        Do(context.Background())
    if err != nil {
        println(err.Error())
    }
    fmt.Printf("update data %s\n", res.Result)
}

//查找
func Gets(params *ElesticInfo) interface{}{
    //通过id查找
    get, err := client.Get().Index(params.Index).Type(params.Type).Id(params.Id).Do(context.Background())
    if err != nil {
        panic(err)
    }
    if get.Found {
        fmt.Printf("Got document %s in version %d from index %s, type %s\n", get.Id, get.Version, get.Index, get.Type)
    }

	return get
}

//搜索
func Query(params *ElesticInfo , queryType string, condition *QueryCondition) interface{}{
    var res *elastic.SearchResult
    var err error

	if queryType == "all" {
		//取所有
		res, err = client.Search(params.Index).Type(params.Type).Do(context.Background())
	} else if queryType == "newQueryStringQuery" {
		//字段相等,"key:value"
		q := elastic.NewQueryStringQuery(condition.Key + ":" + condition.Value)
		res, err = client.Search(params.Index).Type(params.Type).Query(q).Do(context.Background())
	}else if queryType == "phraseQuery" {
		//短语搜索 搜索key字段中有 value
		matchPhraseQuery := elastic.NewMatchPhraseQuery(condition.Key, condition.Value)
		res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	}
  

	if err != nil {
		println(err.Error())
	}
	return res
}

//简单分页
func List(size, page int,params *ElesticInfo) interface{}{
    if size < 0 || page < 1 {
        fmt.Printf("param error")
    }
    res, err := client.Search(params.Index).
        Type(params.Type).
        Size(size).
        From((page - 1) * size).
        Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return res
}



