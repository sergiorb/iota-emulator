package json

import "net/http"
import "fmt"
import "log"
import "io/ioutil"
import "github.com/sergiorb/iota-emulator/src/device"
import "github.com/sergiorb/iota-emulator/src/transports"
import "github.com/sergiorb/iota-emulator/src/agents"
import config "github.com/sergiorb/iota-emulator/src/config/iota-json/v1"

type JsonAgent struct {

	Config	config.IotaJsonV1
}

func (a *JsonAgent) Send(req agents.Request, transport transports.Transport) {

	var request *http.Request;
	var response *http.Response;
	
	if req.Key == "" {
		req.Key = a.Config.Defaults.Key
	}

	if req.Resource == "" {
		req.Resource = a.Config.Defaults.Resource
	}
	
	url := fmt.Sprintf("%v://%v:%v%v?k=%v&i=%v", 
		a.Config.Http.Schema,
		a.Config.Http.Host,
		a.Config.Http.Port,
		req.Resource,
		req.Key,
		req.DeviceId);

	request, err := http.NewRequest("POST", url, nil)
	
    if err == nil {

		request.Header.Add("Content-Type", "application/json");
				
        response, err = (&http.Client{}).Do(request);
	}

	if err != nil {
        log.Printf("ERROR: %s", err)
    }

	if (response != nil && response.StatusCode != 200) {

		defer response.Body.Close()

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		
    	log.Printf(bodyString)
	}

}

func (a *JsonAgent) sendThroughHttp(ttrValue device.AttributeValue) {

}

func (a *JsonAgent) sendThroughMqtt(ttrValue device.AttributeValue) {
	
}

func (a *JsonAgent) sendThroughAmpq(ttrValue device.AttributeValue) {
	
}
