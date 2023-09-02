package main

import (
	"encoding/json"
	"fmt"
)

type District struct {
	ID         string `json:"id"`
	DivisionID string `json:"division_id"`
	Name       string `json:"name"`
	BnName     string `json:"bn_name"`
	Lat        string `json:"lat"`
	Lon        string `json:"lon"`
	URL        string `json:"url"`
}

func main() {
	jsonData := `[
{"id":"1","division_id":"1","name":"Comilla","bn_name":"কুমিল্লা","lat":"23.4682747","lon":"91.1788135","url":"www.comilla.gov.bd"},
{"id":"2","division_id":"1","name":"Feni","bn_name":"ফেনী","lat":"23.023231","lon":"91.3840844","url":"www.feni.gov.bd"},
{"id":"3","division_id":"1","name":"Brahmanbaria","bn_name":"ব্রাহ্মণবাড়িয়া","lat":"23.9570904","lon":"91.1119286","url":"www.brahmanbaria.gov.bd"},
{"id":"4","division_id":"1","name":"Rangamati","bn_name":"রাঙ্গামাটি","lat":"22.65561018","lon":"92.17541121","url":"www.rangamati.gov.bd"},
{"id":"5","division_id":"1","name":"Noakhali","bn_name":"নোয়াখালী","lat":"22.869563","lon":"91.099398","url":"www.noakhali.gov.bd"},
{"id":"6","division_id":"1","name":"Chandpur","bn_name":"চাঁদপুর","lat":"23.2332585","lon":"90.6712912","url":"www.chandpur.gov.bd"},
{"id":"7","division_id":"1","name":"Lakshmipur","bn_name":"লক্ষ্মীপুর","lat":"22.942477","lon":"90.841184","url":"www.lakshmipur.gov.bd"},
{"id":"8","division_id":"1","name":"Chattogram","bn_name":"চট্টগ্রাম","lat":"22.335109","lon":"91.834073","url":"www.chittagong.gov.bd"},
{"id":"9","division_id":"1","name":"Coxsbazar","bn_name":"কক্সবাজার","lat":"21.44315751","lon":"91.97381741","url":"www.coxsbazar.gov.bd"},
{"id":"10","division_id":"1","name":"Khagrachhari","bn_name":"খাগড়াছড়ি","lat":"23.119285","lon":"91.984663","url":"www.khagrachhari.gov.bd"},
{"id":"11","division_id":"1","name":"Bandarban","bn_name":"বান্দরবান","lat":"22.1953275","lon":"92.2183773","url":"www.bandarban.gov.bd"},
{"id":"12","division_id":"2","name":"Sirajganj","bn_name":"সিরাজগঞ্জ","lat":"24.4533978","lon":"89.7006815","url":"www.sirajganj.gov.bd"},
{"id":"13","division_id":"2","name":"Pabna","bn_name":"পাবনা","lat":"23.998524","lon":"89.233645","url":"www.pabna.gov.bd"},
{"id":"14","division_id":"2","name":"Bogura","bn_name":"বগুড়া","lat":"24.8465228","lon":"89.377755","url":"www.bogra.gov.bd"},
{"id":"15","division_id":"2","name":"Rajshahi","bn_name":"রাজশাহী","lat":"24.37230298","lon":"88.56307623","url":"www.rajshahi.gov.bd"},
{"id":"16","division_id":"2","name":"Natore","bn_name":"নাটোর","lat":"24.420556","lon":"89.000282","url":"www.natore.gov.bd"},
{"id":"17","division_id":"2","name":"Joypurhat","bn_name":"জয়পুরহাট","lat":"25.09636876","lon":"89.04004280","url":"www.joypurhat.gov.bd"},
{"id":"18","division_id":"2","name":"Chapainawabganj","bn_name":"চাঁপাইনবাবগঞ্জ","lat":"24.5965034","lon":"88.2775122","url":"www.chapainawabganj.gov.bd"},
{"id":"19","division_id":"2","name":"Naogaon","bn_name":"নওগাঁ","lat":"24.83256191","lon":"88.92485205","url":"www.naogaon.gov.bd"},
{"id":"20","division_id":"3","name":"Jashore","bn_name":"যশোর","lat":"23.16643","lon":"89.2081126","url":"www.jessore.gov.bd"},
{"id":"21","division_id":"3","name":"Satkhira","bn_name":"সাতক্ষীরা","lat":"22.7180905","lon":"89.0687033","url":"www.satkhira.gov.bd"},
{"id":"22","division_id":"3","name":"Meherpur","bn_name":"মেহেরপুর","lat":"23.762213","lon":"88.631821","url":"www.meherpur.gov.bd"},
{"id":"23","division_id":"3","name":"Narail","bn_name":"নড়াইল","lat":"23.172534","lon":"89.512672","url":"www.narail.gov.bd"},
{"id":"24","division_id":"3","name":"Chuadanga","bn_name":"চুয়াডাঙ্গা","lat":"23.6401961","lon":"88.841841","url":"www.chuadanga.gov.bd"},
{"id":"25","division_id":"3","name":"Kushtia","bn_name":"কুষ্টিয়া","lat":"23.901258","lon":"89.120482","url":"www.kushtia.gov.bd"},
{"id":"26","division_id":"3","name":"Magura","bn_name":"মাগুরা","lat":"23.487337","lon":"89.419956","url":"www.magura.gov.bd"},
{"id":"27","division_id":"3","name":"Khulna","bn_name":"খুলনা","lat":"22.815774","lon":"89.568679","url":"www.khulna.gov.bd"},
{"id":"28","division_id":"3","name":"Bagerhat","bn_name":"বাগেরহাট","lat":"22.651568","lon":"89.785938","url":"www.bagerhat.gov.bd"},
{"id":"29","division_id":"3","name":"Jhenaidah","bn_name":"ঝিনাইদহ","lat":"23.5448176","lon":"89.1539213","url":"www.jhenaidah.gov.bd"},
{"id":"30","division_id":"4","name":"Jhalakathi","bn_name":"ঝালকাঠি","lat":"22.6422689","lon":"90.2003932","url":"www.jhalakathi.gov.bd"},
{"id":"31","division_id":"4","name":"Patuakhali","bn_name":"পটুয়াখালী","lat":"22.3596316","lon":"90.3298712","url":"www.patuakhali.gov.bd"},
{"id":"32","division_id":"4","name":"Pirojpur","bn_name":"পিরোজপুর","lat":"22.5781398","lon":"89.9983909","url":"www.pirojpur.gov.bd"},
{"id":"33","division_id":"4","name":"Barisal","bn_name":"বরিশাল","lat":"22.7004179","lon":"90.3731568","url":"www.barisal.gov.bd"},
{"id":"34","division_id":"4","name":"Bhola","bn_name":"ভোলা","lat":"22.685923","lon":"90.648179","url":"www.bhola.gov.bd"},
{"id":"35","division_id":"4","name":"Barguna","bn_name":"বরগুনা","lat":"22.159182","lon":"90.125581","url":"www.barguna.gov.bd"},
{"id":"36","division_id":"5","name":"Sylhet","bn_name":"সিলেট","lat":"24.8897956","lon":"91.8697894","url":"www.sylhet.gov.bd"},
{"id":"37","division_id":"5","name":"Moulvibazar","bn_name":"মৌলভীবাজার","lat":"24.482934","lon":"91.777417","url":"www.moulvibazar.gov.bd"},
{"id":"38","division_id":"5","name":"Habiganj","bn_name":"হবিগঞ্জ","lat":"24.374945","lon":"91.41553","url":"www.habiganj.gov.bd"},
{"id":"39","division_id":"5","name":"Sunamganj","bn_name":"সুনামগঞ্জ","lat":"25.0658042","lon":"91.3950115","url":"www.sunamganj.gov.bd"},
{"id":"40","division_id":"6","name":"Narsingdi","bn_name":"নরসিংদী","lat":"23.932233","lon":"90.71541","url":"www.narsingdi.gov.bd"},
{"id":"41","division_id":"6","name":"Gazipur","bn_name":"গাজীপুর","lat":"24.0022858","lon":"90.4264283","url":"www.gazipur.gov.bd"},
{"id":"42","division_id":"6","name":"Shariatpur","bn_name":"শরীয়তপুর","lat":"23.2060195","lon":"90.3477725","url":"www.shariatpur.gov.bd"},
{"id":"43","division_id":"6","name":"Narayanganj","bn_name":"নারায়ণগঞ্জ","lat":"23.63366","lon":"90.496482","url":"www.narayanganj.gov.bd"},
{"id":"44","division_id":"6","name":"Tangail","bn_name":"টাঙ্গাইল","lat":"24.264145","lon":"89.918029","url":"www.tangail.gov.bd"},
{"id":"45","division_id":"6","name":"Kishoreganj","bn_name":"কিশোরগঞ্জ","lat":"24.444937","lon":"90.776575","url":"www.kishoreganj.gov.bd"},
{"id":"46","division_id":"6","name":"Manikganj","bn_name":"মানিকগঞ্জ","lat":"23.8602262","lon":"90.0018293","url":"www.manikganj.gov.bd"},
{"id":"47","division_id":"6","name":"Dhaka","bn_name":"ঢাকা","lat":"23.7115253","lon":"90.4111451","url":"www.dhaka.gov.bd"},
{"id":"48","division_id":"6","name":"Munshiganj","bn_name":"মুন্সিগঞ্জ","lat":"23.5435742","lon":"90.5354327","url":"www.munshiganj.gov.bd"},
{"id":"49","division_id":"6","name":"Rajbari","bn_name":"রাজবাড়ী","lat":"23.7574305","lon":"89.6444665","url":"www.rajbari.gov.bd"},
{"id":"50","division_id":"6","name":"Madaripur","bn_name":"মাদারীপুর","lat":"23.164102","lon":"90.1896805","url":"www.madaripur.gov.bd"},
{"id":"51","division_id":"6","name":"Gopalganj","bn_name":"গোপালগঞ্জ","lat":"23.0050857","lon":"89.8266059","url":"www.gopalganj.gov.bd"},
{"id":"52","division_id":"6","name":"Faridpur","bn_name":"ফরিদপুর","lat":"23.6070822","lon":"89.8429406","url":"www.faridpur.gov.bd"},
{"id":"53","division_id":"7","name":"Panchagarh","bn_name":"পঞ্চগড়","lat":"26.3411","lon":"88.5541606","url":"www.panchagarh.gov.bd"},
{"id":"54","division_id":"7","name":"Dinajpur","bn_name":"দিনাজপুর","lat":"25.6217061","lon":"88.6354504","url":"www.dinajpur.gov.bd"},
{"id":"55","division_id":"7","name":"Lalmonirhat","bn_name":"লালমনিরহাট","lat":"25.9165451","lon":"89.4532409","url":"www.lalmonirhat.gov.bd"},
{"id":"56","division_id":"7","name":"Nilphamari","bn_name":"নীলফামারী","lat":"25.931794","lon":"88.856006","url":"www.nilphamari.gov.bd"},
{"id":"57","division_id":"7","name":"Gaibandha","bn_name":"গাইবান্ধা","lat":"25.328751","lon":"89.528088","url":"www.gaibandha.gov.bd"},
{"id":"58","division_id":"7","name":"Thakurgaon","bn_name":"ঠাকুরগাঁও","lat":"26.0336945","lon":"88.4616834","url":"www.thakurgaon.gov.bd"},
{"id":"59","division_id":"7","name":"Rangpur","bn_name":"রংপুর","lat":"25.7558096","lon":"89.244462","url":"www.rangpur.gov.bd"},
{"id":"60","division_id":"7","name":"Kurigram","bn_name":"কুড়িগ্রাম","lat":"25.805445","lon":"89.636174","url":"www.kurigram.gov.bd"},
{"id":"61","division_id":"8","name":"Sherpur","bn_name":"শেরপুর","lat":"25.0204933","lon":"90.0152966","url":"www.sherpur.gov.bd"},
{"id":"62","division_id":"8","name":"Mymensingh","bn_name":"ময়মনসিংহ","lat":"24.7465670","lon":"90.4072093","url":"www.mymensingh.gov.bd"},
{"id":"63","division_id":"8","name":"Jamalpur","bn_name":"জামালপুর","lat":"24.937533","lon":"89.937775","url":"www.jamalpur.gov.bd"},
{"id":"64","division_id":"8","name":"Netrokona","bn_name":"নেত্রকোণা","lat":"24.870955","lon":"90.727887","url":"www.netrokona.gov.bd"}
]`

	var districts []District
	err := json.Unmarshal([]byte(jsonData), &districts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(districts)
}
