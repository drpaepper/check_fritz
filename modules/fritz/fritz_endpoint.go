package fritz

type endpoint struct {
	url          string
	service      string
	serviceIndex string
	action       string
}

var WLAN1Associations = endpoint{
	url:          "/upnp/control/wlanconfig1",
	service:      "WLANConfiguration",
	serviceIndex: "1",
	action:       "GetTotalAssociations",
}

var WANPPPConnectionInfo = endpoint{
	url:          "/upnp/control/wanpppconn1",
	service:      "WANPPPConnection",
	serviceIndex: "1",
	action:       "GetInfo",
}

var WANIPConnectionInfo = endpoint{
	url:          "/upnp/control/wanipconnection1",
	service:      "WanIPConnection",
	serviceIndex: "1",
	action:       "GetInfo",
}

var DeviceInfo = endpoint{
	url:          "/upnp/control/deviceinfo",
	service:      "DeviceInfo",
	serviceIndex: "1",
	action:       "GetInfo",
}

var UserInterfaceInfo = endpoint{
	url:          "/upnp/control/userif",
	service:      "UserInterface",
	serviceIndex: "1",
	action:       "GetInfo",
}

var DSLConnectionInfo = endpoint{
	url:          "/upnp/control/wandslifconfig1",
	service:      "WANDSLInterfaceConfig",
	serviceIndex: "1",
	action:       "GetInfo",
}

var WANCommonLinkProperties = endpoint{
	url:          "/upnp/control/wancommonifconfig1",
	service:      "WANCommonInterfaceConfig",
	serviceIndex: "1",
	action:       "GetCommonLinkProperties",
}

var WANCommonOnlineMonitor = endpoint{
	url:          "/upnp/control/wancommonifconfig1",
	service:      "WANCommonInterfaceConfig",
	serviceIndex: "1",
	action:       "X_AVM-DE_GetOnlineMonitor",
}

var HomeAutoDeviceInfo = endpoint{
	url:          "/upnp/control/x_homeauto",
	service:      "X_AVM-DE_Homeauto",
	serviceIndex: "1",
	action:       "GetSpecificDeviceInfos",
}
