package lookup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type APIResponse struct {
	Count          int64               `json:"Count"`
	Message        string              `json:"Message"`
	SearchCriteria string              `json:"SearchCriteria"`
	Results        []APIResponseResult `json:"Results"`
}

type APIResponseResult struct {
	Abs                                 string `json:"ABS"`
	ActiveSafetySysNote                 string `json:"ActiveSafetySysNote"`
	AdaptiveCruiseControl               string `json:"AdaptiveCruiseControl"`
	AdaptiveDrivingBeam                 string `json:"AdaptiveDrivingBeam"`
	AdaptiveHeadlights                  string `json:"AdaptiveHeadlights"`
	AdditionalErrorText                 string `json:"AdditionalErrorText"`
	AirBagLOCCurtain                    string `json:"AirBagLocCurtain"`
	AirBagLOCFront                      string `json:"AirBagLocFront"`
	AirBagLOCKnee                       string `json:"AirBagLocKnee"`
	AirBagLOCSeatCushion                string `json:"AirBagLocSeatCushion"`
	AirBagLOCSide                       string `json:"AirBagLocSide"`
	AutoReverseSystem                   string `json:"AutoReverseSystem"`
	AutomaticPedestrianAlertingSound    string `json:"AutomaticPedestrianAlertingSound"`
	AxleConfiguration                   string `json:"AxleConfiguration"`
	Axles                               string `json:"Axles"`
	BasePrice                           string `json:"BasePrice"`
	BatteryA                            string `json:"BatteryA"`
	BatteryATo                          string `json:"BatteryA_to"`
	BatteryCells                        string `json:"BatteryCells"`
	BatteryInfo                         string `json:"BatteryInfo"`
	BatteryKWh                          string `json:"BatteryKWh"`
	BatteryKWhTo                        string `json:"BatteryKWh_to"`
	BatteryModules                      string `json:"BatteryModules"`
	BatteryPacks                        string `json:"BatteryPacks"`
	BatteryType                         string `json:"BatteryType"`
	BatteryV                            string `json:"BatteryV"`
	BatteryVTo                          string `json:"BatteryV_to"`
	BedLengthIN                         string `json:"BedLengthIN"`
	BedType                             string `json:"BedType"`
	BlindSpotIntervention               string `json:"BlindSpotIntervention"`
	BlindSpotMon                        string `json:"BlindSpotMon"`
	BodyCabType                         string `json:"BodyCabType"`
	BodyClass                           string `json:"BodyClass"`
	BrakeSystemDesc                     string `json:"BrakeSystemDesc"`
	BrakeSystemType                     string `json:"BrakeSystemType"`
	BusFloorConfigType                  string `json:"BusFloorConfigType"`
	BusLength                           string `json:"BusLength"`
	BusType                             string `json:"BusType"`
	CanAacn                             string `json:"CAN_AACN"`
	Cib                                 string `json:"CIB"`
	CashForClunkers                     string `json:"CashForClunkers"`
	ChargerLevel                        string `json:"ChargerLevel"`
	ChargerPowerKW                      string `json:"ChargerPowerKW"`
	CoolingType                         string `json:"CoolingType"`
	CurbWeightLB                        string `json:"CurbWeightLB"`
	CustomMotorcycleType                string `json:"CustomMotorcycleType"`
	DaytimeRunningLight                 string `json:"DaytimeRunningLight"`
	DestinationMarket                   string `json:"DestinationMarket"`
	DisplacementCC                      string `json:"DisplacementCC"`
	DisplacementCI                      string `json:"DisplacementCI"`
	DisplacementL                       string `json:"DisplacementL"`
	Doors                               string `json:"Doors"`
	DriveType                           string `json:"DriveType"`
	DriverAssist                        string `json:"DriverAssist"`
	DynamicBrakeSupport                 string `json:"DynamicBrakeSupport"`
	Edr                                 string `json:"EDR"`
	Esc                                 string `json:"ESC"`
	EVDriveUnit                         string `json:"EVDriveUnit"`
	ElectrificationLevel                string `json:"ElectrificationLevel"`
	EngineConfiguration                 string `json:"EngineConfiguration"`
	EngineCycles                        string `json:"EngineCycles"`
	EngineCylinders                     string `json:"EngineCylinders"`
	EngineHP                            string `json:"EngineHP"`
	EngineHPTo                          string `json:"EngineHP_to"`
	EngineKW                            string `json:"EngineKW"`
	EngineManufacturer                  string `json:"EngineManufacturer"`
	EngineModel                         string `json:"EngineModel"`
	EntertainmentSystem                 string `json:"EntertainmentSystem"`
	ErrorCode                           string `json:"ErrorCode"`
	ErrorText                           string `json:"ErrorText"`
	ForwardCollisionWarning             string `json:"ForwardCollisionWarning"`
	FuelInjectionType                   string `json:"FuelInjectionType"`
	FuelTypePrimary                     string `json:"FuelTypePrimary"`
	FuelTypeSecondary                   string `json:"FuelTypeSecondary"`
	Gcwr                                string `json:"GCWR"`
	GCWRTo                              string `json:"GCWR_to"`
	Gvwr                                string `json:"GVWR"`
	GVWRTo                              string `json:"GVWR_to"`
	KeylessIgnition                     string `json:"KeylessIgnition"`
	LaneCenteringAssistance             string `json:"LaneCenteringAssistance"`
	LaneDepartureWarning                string `json:"LaneDepartureWarning"`
	LaneKeepSystem                      string `json:"LaneKeepSystem"`
	LowerBeamHeadlampLightSource        string `json:"LowerBeamHeadlampLightSource"`
	Make                                string `json:"Make"`
	MakeID                              string `json:"MakeID"`
	Manufacturer                        string `json:"Manufacturer"`
	ManufacturerID                      string `json:"ManufacturerId"`
	Model                               string `json:"Model"`
	ModelID                             string `json:"ModelID"`
	ModelYear                           string `json:"ModelYear"`
	MotorcycleChassisType               string `json:"MotorcycleChassisType"`
	MotorcycleSuspensionType            string `json:"MotorcycleSuspensionType"`
	NCSABodyType                        string `json:"NCSABodyType"`
	NCSAMake                            string `json:"NCSAMake"`
	NCSAMapExcApprovedBy                string `json:"NCSAMapExcApprovedBy"`
	NCSAMapExcApprovedOn                string `json:"NCSAMapExcApprovedOn"`
	NCSAMappingException                string `json:"NCSAMappingException"`
	NCSAModel                           string `json:"NCSAModel"`
	NCSANote                            string `json:"NCSANote"`
	NonLandUse                          string `json:"NonLandUse"`
	Note                                string `json:"Note"`
	OtherBusInfo                        string `json:"OtherBusInfo"`
	OtherEngineInfo                     string `json:"OtherEngineInfo"`
	OtherMotorcycleInfo                 string `json:"OtherMotorcycleInfo"`
	OtherRestraintSystemInfo            string `json:"OtherRestraintSystemInfo"`
	OtherTrailerInfo                    string `json:"OtherTrailerInfo"`
	ParkAssist                          string `json:"ParkAssist"`
	PedestrianAutomaticEmergencyBraking string `json:"PedestrianAutomaticEmergencyBraking"`
	PlantCity                           string `json:"PlantCity"`
	PlantCompanyName                    string `json:"PlantCompanyName"`
	PlantCountry                        string `json:"PlantCountry"`
	PlantState                          string `json:"PlantState"`
	PossibleValues                      string `json:"PossibleValues"`
	Pretensioner                        string `json:"Pretensioner"`
	RearAutomaticEmergencyBraking       string `json:"RearAutomaticEmergencyBraking"`
	RearCrossTrafficAlert               string `json:"RearCrossTrafficAlert"`
	RearVisibilitySystem                string `json:"RearVisibilitySystem"`
	SAEAutomationLevel                  string `json:"SAEAutomationLevel"`
	SAEAutomationLevelTo                string `json:"SAEAutomationLevel_to"`
	SeatBeltsAll                        string `json:"SeatBeltsAll"`
	SeatRows                            string `json:"SeatRows"`
	Seats                               string `json:"Seats"`
	SemiautomaticHeadlampBeamSwitching  string `json:"SemiautomaticHeadlampBeamSwitching"`
	Series                              string `json:"Series"`
	Series2                             string `json:"Series2"`
	SteeringLocation                    string `json:"SteeringLocation"`
	SuggestedVIN                        string `json:"SuggestedVIN"`
	Tpms                                string `json:"TPMS"`
	TopSpeedMPH                         string `json:"TopSpeedMPH"`
	TrackWidth                          string `json:"TrackWidth"`
	TractionControl                     string `json:"TractionControl"`
	TrailerBodyType                     string `json:"TrailerBodyType"`
	TrailerLength                       string `json:"TrailerLength"`
	TrailerType                         string `json:"TrailerType"`
	TransmissionSpeeds                  string `json:"TransmissionSpeeds"`
	TransmissionStyle                   string `json:"TransmissionStyle"`
	Trim                                string `json:"Trim"`
	Trim2                               string `json:"Trim2"`
	Turbo                               string `json:"Turbo"`
	Vin                                 string `json:"VIN"`
	ValveTrainDesign                    string `json:"ValveTrainDesign"`
	VehicleDescriptor                   string `json:"VehicleDescriptor"`
	VehicleType                         string `json:"VehicleType"`
	WheelBaseLong                       string `json:"WheelBaseLong"`
	WheelBaseShort                      string `json:"WheelBaseShort"`
	WheelBaseType                       string `json:"WheelBaseType"`
	WheelSizeFront                      string `json:"WheelSizeFront"`
	WheelSizeRear                       string `json:"WheelSizeRear"`
	Wheels                              string `json:"Wheels"`
	Windows                             string `json:"Windows"`
}

type VinInfo struct {
	Vin    string
	Year   string
	Result *APIResponseResult
}

var apiURL = "https://vpic.nhtsa.dot.gov/api/vehicles/decodevinvalues/%s?format=json&modelyear=%s"

func RequestVinInfo(data *[]VinInfo) error {
	g := new(errgroup.Group)

	for vinIdx, vinInfo := range *data {
		vinIdx := vinIdx
		vinInfo := vinInfo

		g.Go(func() error {
			requestURL := fmt.Sprintf(apiURL, vinInfo.Vin, vinInfo.Year)

			req, err := http.NewRequest(http.MethodGet, requestURL, nil)
			if err != nil {
				return fmt.Errorf("error creating request: %s", err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return fmt.Errorf("error making HTTP request: %s", err)
			}

			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return fmt.Errorf("error reading response body: %s", err)
			}

			var apiResponse APIResponse

			jsonErr := json.Unmarshal(resBody, &apiResponse)
			if jsonErr != nil {
				return fmt.Errorf("error unmarshalling: %s", jsonErr)
			}

			(*data)[vinIdx].Result = &apiResponse.Results[0]

			return nil
		})

		time.Sleep(100 * time.Millisecond)
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error performing requests. %s", err)
	}

	return nil
}
