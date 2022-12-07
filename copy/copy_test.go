package json
import (
   "bytes"
   "encoding/gob"
   "encoding/json"
   "fmt"
   "testing"
)
/**
 * @Description: 请求信息
 */
type BidRequest struct {
   ID     string  `json:"id"`
   Imp    []*Imp  `json:"imp"`
   Device *Device `json:"device"`
}
/**
 * @Description: imp对象
 */
type Imp struct {
   ID          string  `json:"id"`
   Tagid       string  `json:"tagid"`
   Bidfloor    float64 `json:"bidfloor"`
}
/**
 * @Description: 设备信息
 */
type Device struct {
   Ua         string `json:"ua"`
   IP         string `json:"ip"`
   Geo        *Geo   `json:"geo"`
   Make       string `json:"make"`
   Model      string `json:"model"`
   Os         string `json:"os"`
   Osv        string `json:"osv"`
}
/**
 * @Description: 地理位置信息
 */
type Geo struct {
   Lat     int    `json:"lat"`
   Lon     int    `json:"lon"`
   Country string `json:"country"`
   Region  string `json:"region"`
   City    string `json:"city"`
}
/**
 * @Description: 利用gob进行深拷贝
 */
func DeepCopyByGob(src,dst interface{}) error {
   var buffer bytes.Buffer
   if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
      return err
   }
   return gob.NewDecoder(&buffer).Decode(dst)
}
/**
 * @Description: 利用json进行深拷贝
 */
func DeepCopyByJson(src,dst *BidRequest) error{
   if tmp, err := json.Marshal(&src);err!=nil{
      return err
   }else {
      err = json.Unmarshal(tmp, dst)
      return err
   }
}

/**
 * @Description: 通过自定义进行copy
 */
func DeepCopyByCustom(src,dst *BidRequest){
   dst.ID=src.ID
   dst.Device=&Device{
      Ua: src.Device.Ua,
      IP: src.Device.IP,
      Geo: &Geo{
         Lat: src.Device.Geo.Lat,
         Lon: src.Device.Geo.Lon,
      },
      Make: src.Device.Make,
      Model: src.Device.Model,
      Os: src.Device.Os,
      Osv: src.Device.Osv,
   }
   dst.Imp=make([]*Imp,len(src.Imp))
   for index,imp:=range src.Imp{
      //注意此处因为imp对象里无指针对象,所以可以直接使用等于
      dst.Imp[index]=imp
   }
}

func initData()*BidRequest  {
   str:="{"id":"MM7dIXz4H05qtmViqnY5dW","imp":[{"id":"1","tagid":"3979722720","bidfloor":0.01}],"device":{"ua":"Mozilla/5.0 (Linux; Android 10; SM-G960N Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.115 Mobile Safari/537.36 (Mobile; afma-sdk-a-v212621039.212621039.0)","ip":"192.168.1.0","geo":{"lat":0,"lon":0,"country":"KOR","region":"KR-11","city":"Seoul"},"make":"samsung","model":"sm-g960n","os":"android","osv":"10"}}"
   ans:=new(BidRequest)
   json.Unmarshal([]byte(str),&ans)
   return ans
}

/**
 * @Description: 压测深拷贝 -gob
 */
func BenchmarkDeepCopy_Gob(b *testing.B)  {
   src:=initData()
   b.ResetTimer()
   for i:=0;i<b.N;i++{
      DeepCopyByGob(src,new(BidRequest))
   }
}

/**
 * @Description: 压测深拷贝 -json
 */
func BenchmarkDeepCopy_Json(b *testing.B)  {
   src:=initData()
   b.ResetTimer()
   for i:=0;i<b.N;i++{
      DeepCopyByJson(src,new(BidRequest))
   }
}
/**
 * @Description: 压测深拷贝 -custom
 */
func BenchmarkDeepCopy_custom(b *testing.B)  {
   src:=initData()
   b.ResetTimer()
   for i:=0;i<b.N;i++{
      DeepCopyByCustom(src,new(BidRequest))
   }
}
/**
 * @Description: 测试拷贝是否ok,go test -bench=. -benchmem
 */
/** 
 * 从性能上来讲 custom>json>gob,
 * 从代码数量上来讲 gob>json>custom ,
 * 因此具体使用时应该充分考虑性能和代码复杂度,
 * 若性能要求不是很高建议gob方法,其比较简洁并且利于生成工具包,
 * 若要求性能则尽量使用custom,此处不偷懒可以提高性能哦。
 * 若是性能要求在中间,则可以使用json先序列化,再反序列化赋值。 
*/
func TestCpoyIsOk(t *testing.T)  {
   src:=initData()
   //1.gob
   dst01:=new(BidRequest)
   DeepCopyByGob(src,dst01)
   bs01, _ := json.Marshal(dst01)
   fmt.Printf("%v\n",string(bs01))
   //2.json
   dst02:=new(BidRequest)
   DeepCopyByJson(src,dst02)
   bs02, _ := json.Marshal(dst02)
   fmt.Printf("%v\n",string(bs02))
   //3.custom
   dst03:=new(BidRequest)
   DeepCopyByCustom(src,dst03)
   bs03, _ := json.Marshal(dst02)
   fmt.Printf("%v\n",string(bs03))
}