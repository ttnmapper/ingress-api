package types

/*
{
   "network_type":"NS_TTN_V2",
   "network_address":"eu.thethings.network",
   "app_id":"jpm_testing",
   "dev_id":"things_uno_jpm",
   "dev_eui":"0004A30B001C684F",
   "time":1610281183625000000,
   "port":1,
   "counter":1,
   "frequency":867700000,
   "modulation":"LORA",
   "bandwidth":125000,
   "spreading_factor":7,
   "coding_rate":"4/5",
   "gateways":[
      {
         "gtw_id":"eui-3133303748005c00",
         "gtw_eui":"3133303748005C00",
         "time":1610281181612000000,
         "timestamp":116210767,
         "channel":6,
         "rssi":-89.0,
         "snr":10.0
      }
   ],
   "latitude":-33.954749178527024,
   "longitude":22.43430212584409,
   "altitude":273.32907678362216,
   "accuracy_meters":24.0,
   "accuracy_source":"gps",
   "userid":"d2883904-9aa7-43f6-b7ee-e9506341291f",
   "useragent":"Android10 App21:2020.12.25"
}
*/

type TtnMapperAndroidMessageV4 = TtnMapperUplinkMessage
