package controller

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go_blog/api/v1"
)

var Dashboart = cDashboart{}

type cDashboart struct {
}

func (c *cDashboart) Dashboart(ctx context.Context, req *v1.DashboartReq) (res *v1.DashboardRes, err error) {
	jsonContent := `{"uv":37,"blogCount":26,"pv":136,"cityVisitor":[{"city":"上海市","uv":2375},{"city":"广州市","uv":1513},{"city":"北京市","uv":1374},{"city":"成都市","uv":1053},{"city":"杭州市","uv":1016},{"city":"郑州市","uv":987},{"city":"深圳市","uv":957},{"city":"武汉市","uv":871},{"city":"南京市","uv":711},{"city":"长沙市","uv":528},{"city":"重庆市","uv":519},{"city":"西安市","uv":396},{"city":"南昌市","uv":347},{"city":"东莞市","uv":328},{"city":"合肥市","uv":327},{"city":"济南市","uv":300},{"city":"福州市","uv":292},{"city":"苏州市","uv":275},{"city":"南宁市","uv":243},{"city":"厦门市","uv":238},{"city":"天津市","uv":228},{"city":"佛山市","uv":200},{"city":"青岛市","uv":186},{"city":"石家庄市","uv":173},{"city":"大连市","uv":171},{"city":"宁波市","uv":166},{"city":"贵阳市","uv":151},{"city":"昆明市","uv":138},{"city":"汕头市","uv":126},{"city":"无锡市","uv":124},{"city":"温州市","uv":123},{"city":"沈阳市","uv":116},{"city":"洛阳市","uv":113},{"city":"哈尔滨市","uv":107},{"city":"南阳市","uv":102},{"city":"常州市","uv":101},{"city":"太原市","uv":89},{"city":"赣州市","uv":89},{"city":"惠州市","uv":88},{"city":"泉州市","uv":87},{"city":"珠海市","uv":87},{"city":"长春市","uv":86},{"city":"衡阳市","uv":85},{"city":"中山市","uv":84},{"city":"桂林市","uv":83},{"city":"阳泉市","uv":79},{"city":"徐州市","uv":78},{"city":"周口市","uv":77},{"city":"湛江市","uv":76},{"city":"潍坊市","uv":74},{"city":"江门市","uv":72},{"city":"海口市","uv":72},{"city":"保定市","uv":70},{"city":"揭阳市","uv":68},{"city":"开封市","uv":67},{"city":"柳州市","uv":67},{"city":"镇江市","uv":67},{"city":"湘潭市","uv":66},{"city":"烟台市","uv":64},{"city":"南通市","uv":62},{"city":"岳阳市","uv":62},{"city":"新乡市","uv":62},{"city":"金华市","uv":60},{"city":"绵阳市","uv":58},{"city":"邯郸市","uv":58},{"city":"临沂市","uv":54},{"city":"台州市","uv":53},{"city":"呼和浩特市","uv":53},{"city":"驻马店市","uv":53},{"city":"九江市","uv":52},{"city":"嘉兴市","uv":51},{"city":"唐山市","uv":49},{"city":"廊坊市","uv":48},{"city":"株洲市","uv":48},{"city":"兰州市","uv":47},{"city":"安阳市","uv":47},{"city":"沧州市","uv":46},{"city":"乌鲁木齐市","uv":44},{"city":"宜春市","uv":44},{"city":"黄冈市","uv":44},{"city":"泰安市","uv":43},{"city":"济宁市","uv":43},{"city":"绍兴市","uv":42},{"city":"邵阳市","uv":42},{"city":"黄石市","uv":41},{"city":"威海市","uv":40},{"city":"渭南市","uv":40},{"city":"盐城市","uv":40},{"city":"上饶市","uv":39},{"city":"信阳市","uv":39},{"city":"扬州市","uv":38},{"city":"吉安市","uv":37},{"city":"平顶山市","uv":37},{"city":"淮安市","uv":37},{"city":"肇庆市","uv":36},{"city":"芜湖市","uv":36},{"city":"阜阳市","uv":36},{"city":"襄阳","uv":35},{"city":"宜宾市","uv":33},{"city":"淮南市","uv":33},{"city":"漳州市","uv":33},{"city":"焦作市","uv":33},{"city":"茂名市","uv":33},{"city":"贵港市","uv":33},{"city":"菏泽市","uv":32},{"city":"安庆市","uv":31},{"city":"崇左市","uv":31},{"city":"遵义市","uv":31},{"city":"银川市","uv":31},{"city":"孝感市","uv":30},{"city":"德阳市","uv":30},{"city":"淄博市","uv":30},{"city":"聊城市","uv":30},{"city":"连云港市","uv":30},{"city":"咸阳市","uv":29},{"city":"宝鸡市","uv":29},{"city":"永州市","uv":29},{"city":"汕尾市","uv":29},{"city":"秦皇岛市","uv":29},{"city":"衢州市","uv":29},{"city":"六安市","uv":28},{"city":"南充市","uv":28},{"city":"商丘市","uv":28},{"city":"梅州市","uv":28},{"city":"荆州市","uv":27},{"city":"马鞍山市","uv":27},{"city":"常德市","uv":26},{"city":"枣庄市","uv":26},{"city":"莆田市","uv":26},{"city":"晋中市","uv":25},{"city":"玉林市","uv":25},{"city":"运城市","uv":25},{"city":"邢台市","uv":25},{"city":"长治市","uv":25},{"city":"泰州市","uv":24},{"city":"泸州市","uv":24},{"city":"滁州市","uv":24},{"city":"蚌埠市","uv":24},{"city":"韶关市","uv":24},{"city":"张家口市","uv":23},{"city":"百色市","uv":23},{"city":"龙岩市","uv":23},{"city":"宿迁市","uv":22},{"city":"自贡市","uv":22},{"city":"衡水市","uv":22},{"city":"许昌市","uv":22},{"city":"宜昌市","uv":21},{"city":"萍乡市","uv":21},{"city":"郴州市","uv":21},{"city":"鄂州市","uv":21},{"city":"张家界市","uv":20},{"city":"河源市","uv":20},{"city":"清远市","uv":20},{"city":"湖州市","uv":20},{"city":"包头市","uv":19},{"city":"北海市","uv":19},{"city":"十堰市","uv":19},{"city":"德州市","uv":19},{"city":"潮州市","uv":19},{"city":"赤峰市","uv":19},{"city":"东营市","uv":18},{"city":"新余市","uv":18},{"city":"达州市","uv":18},{"city":"钦州市","uv":18},{"city":"南平市","uv":17},{"city":"怀化市","uv":17},{"city":"抚州市","uv":17},{"city":"日照市","uv":17},{"city":"毕节","uv":17},{"city":"湘西","uv":17},{"city":"黔南","uv":17},{"city":"乐山市","uv":16},{"city":"大同市","uv":16},{"city":"宿州市","uv":16},{"city":"巴中市","uv":16},{"city":"忻州市","uv":16},{"city":"营口市","uv":16},{"city":"吕梁市","uv":15},{"city":"宁德市","uv":15},{"city":"景德镇市","uv":15},{"city":"来宾市","uv":15},{"city":"梧州市","uv":15},{"city":"榆林市","uv":15},{"city":"荆门市","uv":15},{"city":"临汾市","uv":14},{"city":"亳州市","uv":14},{"city":"吉林市","uv":14},{"city":"宣城市","uv":14},{"city":"漯河市","uv":14},{"city":"云浮市","uv":13},{"city":"咸宁市","uv":13},{"city":"娄底市","uv":13},{"city":"濮阳市","uv":13},{"city":"益阳市","uv":13},{"city":"绥化市","uv":13},{"city":"舟山市","uv":13},{"city":"鹤壁市","uv":13},{"city":"三明市","uv":12},{"city":"承德市","uv":12},{"city":"汉中市","uv":12},{"city":"眉山市","uv":12},{"city":"内江市","uv":11},{"city":"彰化县","uv":11},{"city":"朝阳市","uv":11},{"city":"葫芦岛市","uv":11},{"city":"遂宁市","uv":11},{"city":"鞍山市","uv":11},{"city":"六盘水市","uv":10},{"city":"广安市","uv":10},{"city":"恩施","uv":10},{"city":"曲靖市","uv":10},{"city":"河池市","uv":10},{"city":"淮北市","uv":10},{"city":"白山市","uv":10},{"city":"莱芜市","uv":10},{"city":"西宁市","uv":10},{"city":"锦州市","uv":10},{"city":"阳江市","uv":10},{"city":"丽水市","uv":9},{"city":"乌兰察布市","uv":9},{"city":"呼伦贝尔市","uv":9},{"city":"安顺市","uv":9},{"city":"巴彦淖尔市","uv":9},{"city":"抚顺市","uv":9},{"city":"攀枝花市","uv":9},{"city":"朔州市","uv":9},{"city":"滨州市","uv":9},{"city":"三门峡市","uv":8},{"city":"延边","uv":8},{"city":"红河","uv":8},{"city":"资阳市","uv":8},{"city":"铜仁","uv":8},{"city":"阿克苏","uv":8},{"city":"三亚市","uv":7},{"city":"大庆市","uv":7},{"city":"安康市","uv":7},{"city":"昌吉","uv":7},{"city":"本溪市","uv":7},{"city":"铁岭市","uv":7},{"city":"阜新市","uv":7},{"city":"黄山市","uv":7},{"city":"黔东南","uv":7},{"city":"中卫市","uv":6},{"city":"克拉玛依市","uv":6},{"city":"兴安盟","uv":6},{"city":"巴音郭楞","uv":6},{"city":"牡丹江市","uv":6},{"city":"甘孜","uv":6},{"city":"石嘴山市","uv":6},{"city":"通辽市","uv":6},{"city":"鄂尔多斯市","uv":6},{"city":"鹰潭市","uv":6},{"city":"齐齐哈尔市","uv":6},{"city":"丹东市","uv":5},{"city":"双鸭山市","uv":5},{"city":"台北","uv":5},{"city":"平凉市","uv":5},{"city":"延安市","uv":5},{"city":"昭通市","uv":5},{"city":"晋城市","uv":5},{"city":"荃湾","uv":5},{"city":"黑河市","uv":5},{"city":"黔西南","uv":5},{"city":"克孜勒苏","uv":4},{"city":"凉山","uv":4},{"city":"商洛市","uv":4},{"city":"四平市","uv":4},{"city":"塔城","uv":4},{"city":"大理","uv":4},{"city":"贺州市","uv":4},{"city":"酒泉市","uv":4},{"city":"铜陵市","uv":4},{"city":"防城港市","uv":4},{"city":"阿勒泰","uv":4},{"city":"鸡西市","uv":4},{"city":"伊犁","uv":3},{"city":"佳木斯市","uv":3},{"city":"保山市","uv":3},{"city":"儋州","uv":3},{"city":"喀什","uv":3},{"city":"固原市","uv":3},{"city":"定西市","uv":3},{"city":"广元市","uv":3},{"city":"庆阳市","uv":3},{"city":"张掖市","uv":3},{"city":"海南","uv":3},{"city":"海西","uv":3},{"city":"白银市","uv":3},{"city":"西双版纳","uv":3},{"city":"阿坝","uv":3},{"city":"随州市","uv":3},{"city":"吐鲁番","uv":2},{"city":"天水市","uv":2},{"city":"德宏","uv":2},{"city":"普洱市","uv":2},{"city":"楚雄","uv":2},{"city":"济源","uv":2},{"city":"玉溪市","uv":2},{"city":"通化市","uv":2},{"city":"铜川市","uv":2},{"city":"锡林郭勒","uv":2},{"city":"陇南市","uv":2},{"city":"雅安市","uv":2},{"city":"临沧市","uv":1},{"city":"乌海市","uv":1},{"city":"吴忠市","uv":1},{"city":"哈密","uv":1},{"city":"嘉峪关市","uv":1},{"city":"拉萨市","uv":1},{"city":"文山","uv":1},{"city":"昌都","uv":1},{"city":"松原市","uv":1},{"city":"武威市","uv":1},{"city":"池州市","uv":1},{"city":"海东","uv":1},{"city":"白城市","uv":1},{"city":"盘锦市","uv":1},{"city":"辽源市","uv":1},{"city":"迪庆","uv":1},{"city":"金昌市","uv":1},{"city":"高雄","uv":1},{"city":"鹤岗市","uv":1}],"tag":{"legend":["密码学","RabbitMQ","消息队列","布隆过滤器","定时任务","Redis","Spring Boot","算法","数据结构","Sublime Text 3","Java","Python 3","Vue Router","Vue","心情随写","DFS","BFS","图像识别","连连看","Typora","PicGo","GitHub","jsDelivr","Swing","五子棋","Python-Flask","Nginx","归并排序","败者树","外部排序","跳表SkipList","KMP算法","字符串"],"series":[{"id":43,"name":"Vue","value":3},{"id":46,"name":"Vue Router","value":1},{"id":42,"name":"心情随写","value":4},{"id":9,"name":"Nginx","value":2},{"id":10,"name":"Python-Flask","value":1},{"id":47,"name":"Python 3","value":7},{"id":49,"name":"Sublime Text 3","value":1},{"id":3,"name":"跳表SkipList","value":1},{"id":50,"name":"数据结构","value":3},{"id":4,"name":"外部排序","value":1},{"id":5,"name":"败者树","value":1},{"id":6,"name":"归并排序","value":1},{"id":17,"name":"连连看","value":1},{"id":19,"name":"BFS","value":1},{"id":18,"name":"图像识别","value":2},{"id":1,"name":"字符串","value":1},{"id":2,"name":"KMP算法","value":1},{"id":51,"name":"算法","value":2},{"id":52,"name":"Spring Boot","value":2},{"id":53,"name":"Redis","value":2},{"id":54,"name":"定时任务","value":1},{"id":11,"name":"五子棋","value":1},{"id":12,"name":"Swing","value":1},{"id":20,"name":"DFS","value":1},{"id":48,"name":"Java","value":3},{"id":55,"name":"布隆过滤器","value":1},{"id":56,"name":"消息队列","value":1},{"id":57,"name":"RabbitMQ","value":1},{"id":13,"name":"jsDelivr","value":1},{"id":14,"name":"GitHub","value":1},{"id":15,"name":"PicGo","value":1},{"id":16,"name":"Typora","value":1},{"id":58,"name":"密码学","value":1}]},"category":{"legend":["学习笔记","个人项目","技术杂烩","心情随写"],"series":[{"id":32,"name":"心情随写","value":4},{"id":33,"name":"技术杂烩","value":8},{"id":34,"name":"个人项目","value":5},{"id":35,"name":"学习笔记","value":9}]},"visitRecord":{"date":["11-09","11-10","11-11","11-12","11-13","11-14","11-15","11-16","11-17","11-18","11-19","11-20","11-21","11-22","11-23","11-24","11-25","11-26","11-27","11-28","11-29","11-30","12-01","12-02","12-03","12-04","12-05","12-06","12-07","12-08"],"uv":[95,93,108,101,81,111,112,87,96,117,92,79,90,83,84,90,90,84,81,91,113,92,74,70,66,81,77,81,78,71],"pv":[450,447,540,416,285,632,475,341,481,599,416,395,502,395,378,406,460,377,457,350,525,398,326,268,199,285,269,353,225,302]},"commentCount":770}`
	j := gjson.New(jsonContent)
	err = gconv.Structs(j, &res)
	if err != nil {
		return nil, err
	}
	return
}
