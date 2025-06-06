package fritz

type Endpoint struct {
	url          string
	service      string
	serviceIndex string
	action       string
}

var WLAN1Associations = Endpoint{
	url:          "/upnp/control/wlanconfig1",
	service:      "WLANConfiguration",
	serviceIndex: "1",
	action:       "GetTotalAssociations",
}

var WANPPPConnectionInfo = Endpoint{
	url:          "/upnp/control/wanpppconn1",
	service:      "WANPPPConnection",
	serviceIndex: "1",
	action:       "GetInfo",
}

var WANIPConnectionInfo = Endpoint{
	url:          "/upnp/control/wanipconnection1",
	service:      "WanIPConnection",
	serviceIndex: "1",
	action:       "GetInfo",
}

var DeviceInfo = Endpoint{
	url:          "/upnp/control/deviceinfo",
	service:      "DeviceInfo",
	serviceIndex: "1",
	action:       "GetInfo",
}

var UserInterfaceInfo = Endpoint{
	url:          "/upnp/control/userif",
	service:      "UserInterface",
	serviceIndex: "1",
	action:       "GetInfo",
}

var DSLConnectionInfo = Endpoint{
	url:          "/upnp/control/wandslifconfig1",
	service:      "WANDSLInterfaceConfig",
	serviceIndex: "1",
	action:       "GetInfo",
}

var WANCommonLinkProperties = Endpoint{
	url:          "/upnp/control/wancommonifconfig1",
	service:      "WANCommonInterfaceConfig",
	serviceIndex: "1",
	action:       "GetCommonLinkProperties",
}

var WANCommonOnlineMonitor = Endpoint{
	url:          "/upnp/control/wancommonifconfig1",
	service:      "WANCommonInterfaceConfig",
	serviceIndex: "1",
	action:       "X_AVM-DE_GetOnlineMonitor",
}

var HomeAutoDeviceInfo = Endpoint{
	url:          "/upnp/control/x_homeauto",
	service:      "X_AVM-DE_Homeauto",
	serviceIndex: "1",
	action:       "GetSpecificDeviceInfos",
}

var WANCommonActiveProvider = Endpoint{
	url:          "/upnp/control/wancommonifconfig1",
	service:      "WANCommonInterfaceConfig",
	serviceIndex: "1",
	action:       "X_AVM-DE_GetActiveProvider",
}

var WLANConfigurationTotalAssociations = Endpoint{
	url:          "/upnp/control/wlanconfig2",
	service:      "WLANConfiguration",
	serviceIndex: "2",
	action:       "GetTotalAssociations",
}
