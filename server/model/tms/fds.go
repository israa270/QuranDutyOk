package tms

// ResponseFDS returned from FDS APIs
type ResponseFDS struct {
	URL     string `json:"url"`
	MD5     string `json:"md5"`
	Path    string `json:"path"`
	Domain  string `json:"domain"`
	Scene   string `json:"scene"`
	Size    int    `json:"size"`
	Mtime   int    `json:"mtime"`
	Scenes  string `json:"scenes"`
	RetMsg  string `json:"retmsg"`
	RetCode int    `json:"retcode"`
	Src     string `json:"src"`
}


// map[string]interface {}{"domain":"http://tms-fds:8080", 
// "md5":"250ef58bf58b426f16e0390f4654cc42", 
// "mtime":1.687771316e+09, 
// "path":"/group1/default/20230626/17/21/5/889db0b8bc1c7937aa50c08d4d3259f4_20230626092124.apk", 
// "retcode":0, "retmsg":"",
//  "scene":"default", "scenes":"default", 
//  "size":1.1417514e+07,
//   "src":"/group1/default/20230626/17/21/5/889db0b8bc1c7937aa50c08d4d3259f4_20230626092124.apk", 
//   "url":"http://tms-fds:8080/group1/default/20230626/17/21/5/889db0b8bc1c7937aa50c08d4d3259f4_20230626092124.apk?name=889db0b8bc1c7937aa50c08d4d3259f4_20230626092124.apk&download=1"}