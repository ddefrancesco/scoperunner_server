package scopeparser

import "fmt"

type Info string
type InfoCommandValue string
type ErrUnknownInfoCommand error

const (
	InfoAltitude          Info = "altitude"
	InfoLTT               Info = "ltt"
	InfoBrighterMagLimit  Info = "browse_bml"
	InfoCurrentDate       Info = "current_date"
	InfoClockFmt12        Info = "clock_12"
	InfoDeclination       Info = "declination"
	InfoClockFmt24        Info = "clock_24"
	InfoSelectedTargetDec Info = "sel_target_dec"
	InfoFieldDiameter     Info = "field_diameter"
	InfoFainterMagLimit   Info = "fainter_mag_limit"
	InfoUTCOffset         Info = "utc_offset"
	InfoCurrentSiteLong   Info = "current_site_long"
	InfoHighLimit         Info = "high_limit"
	InfoLocalTime24h      Info = "local_time_24h"
	InfoLargerSizeLimit   Info = "larger_size_limit"
	InfoLowerSizeLimit    Info = "lower_size_limit"
	InfoMinimumQuality    Info = "minimum_find_quality"
	InfoRA                Info = "ra"
	InfoCurrentTargetRA   Info = "sel_target_ra"
	InfoSiderealTime      Info = "sidereal_time"
	InfoSmallerSizeLimit  Info = "smaller_size_limit"
	InfoTrackingRate      Info = "tracking_rate"
	InfoCurrentSiteLat    Info = "current_site_lat"
	InfoFirmwareDate      Info = "firmware_date"
	InfoFirmwareVersion   Info = "firmware_version"
	InfoProductName       Info = "product_name"
	InfoFirmwareTime      Info = "firmware_time"
	InfoDeepsky           Info = "deepsky"
	InfoAzimuth           Info = "azimuth"
)

const (
	InfoAltitudeCmd          InfoCommandValue = ":GA#"
	InfoLTTCmd               InfoCommandValue = ":Ga#"
	InfoBrighterMagLimitCmd  InfoCommandValue = ":Gb#"
	InfoCurrentDateCmd       InfoCommandValue = ":GC#"
	InfoClockFmt12Cmd        InfoCommandValue = ":Gc#"
	InfoDeclinationCmd       InfoCommandValue = ":GD#"
	InfoSelectedTargetDecCmd InfoCommandValue = ":Gd#"
	InfoFieldDiameterCmd     InfoCommandValue = ":GF#"
	InfoFainterMagLimitCmd   InfoCommandValue = ":Gf#"
	InfoUTCOffsetCmd         InfoCommandValue = ":GG#"
	InfoCurrentSiteLongCmd   InfoCommandValue = ":Gg#"
	InfoHighLimitCmd         InfoCommandValue = ":Gh#"
	InfoLocalTime24hCmd      InfoCommandValue = ":GL#"
	InfoLargerSizeLimitCmd   InfoCommandValue = ":GI#"
	InfoLowerSizeLimitCmd    InfoCommandValue = ":Go#"
	InfoMinimumQualityCmd    InfoCommandValue = ":Gq#"
	InfoRACmd                InfoCommandValue = ":GR#"
	InfoCurrentTargetRACmd   InfoCommandValue = ":Gr#"
	InfoSiderealTimeCmd      InfoCommandValue = ":GS#"
	InfoSmallerSizeLimitCmd  InfoCommandValue = ":Gs#"
	InfoTrackingRateCmd      InfoCommandValue = ":GT#"
	InfoCurrentSiteLatCmd    InfoCommandValue = ":Gt#"
	InfoFirmwareDateCmd      InfoCommandValue = ":GVD#"
	InfoFirmwareVersionCmd   InfoCommandValue = ":GVN#"
	InfoProductNameCmd       InfoCommandValue = ":GVP#"
	InfoFirmwareTimeCmd      InfoCommandValue = ":GVT#"
	InfoDeepskyCmd           InfoCommandValue = ":Gy#"
	InfoAzimuthCmd           InfoCommandValue = ":GZ#"
)

type InfoCommand struct {
	Info  Info
	Value InfoCommandValue
	Err   ErrUnknownInfoCommand
}

func NewInfoCommand(m Info) *InfoCommand {
	infoCommand := &InfoCommand{
		Info: m,
	}
	return infoCommand
}

func (ic *InfoCommand) InitMap() map[Info]InfoCommandValue {
	infoMap := make(map[Info]InfoCommandValue)
	infoMap[InfoAltitude] = InfoAltitudeCmd
	infoMap[InfoLTT] = InfoLTTCmd
	infoMap[InfoBrighterMagLimit] = InfoBrighterMagLimitCmd
	infoMap[InfoCurrentDate] = InfoCurrentDateCmd
	infoMap[InfoClockFmt12] = InfoClockFmt12Cmd
	infoMap[InfoDeclination] = InfoDeclinationCmd
	infoMap[InfoSelectedTargetDec] = InfoSelectedTargetDecCmd
	infoMap[InfoFieldDiameter] = InfoFieldDiameterCmd
	infoMap[InfoFainterMagLimit] = InfoFainterMagLimitCmd
	infoMap[InfoUTCOffset] = InfoUTCOffsetCmd
	infoMap[InfoCurrentSiteLong] = InfoCurrentSiteLongCmd
	infoMap[InfoHighLimit] = InfoHighLimitCmd
	infoMap[InfoLocalTime24h] = InfoLocalTime24hCmd
	infoMap[InfoLargerSizeLimit] = InfoLargerSizeLimitCmd
	infoMap[InfoLowerSizeLimit] = InfoLowerSizeLimitCmd
	infoMap[InfoMinimumQuality] = InfoMinimumQualityCmd
	infoMap[InfoRA] = InfoRACmd
	infoMap[InfoCurrentTargetRA] = InfoCurrentTargetRACmd
	infoMap[InfoSiderealTime] = InfoSiderealTimeCmd
	infoMap[InfoSmallerSizeLimit] = InfoSmallerSizeLimitCmd
	infoMap[InfoTrackingRate] = InfoTrackingRateCmd
	infoMap[InfoCurrentSiteLat] = InfoCurrentSiteLatCmd
	infoMap[InfoFirmwareDate] = InfoFirmwareDateCmd
	infoMap[InfoFirmwareVersion] = InfoFirmwareVersionCmd
	infoMap[InfoProductName] = InfoProductNameCmd
	infoMap[InfoFirmwareTime] = InfoFirmwareTimeCmd
	infoMap[InfoDeepsky] = InfoDeepskyCmd
	infoMap[InfoAzimuth] = InfoAzimuthCmd
	return infoMap
}

func (ic *InfoCommand) ParseMap() (InfoCommandValue, error) {

	aMap := ic.InitMap()
	ic.Value = aMap[ic.Info]
	if _, ok := aMap[ic.Info]; ok {
		return aMap[ic.Info], nil
	}
	return "error", fmt.Errorf("unknown info")

}

func (ic *InfoCommand) String() string {
	return string(ic.Value)
}

func (ic *InfoCommand) StringValue() string {
	return string(ic.Value)
}

func (ic *InfoCommand) Error() string {
	return fmt.Sprintf("unknown info command: %s", ic.Err)
}
