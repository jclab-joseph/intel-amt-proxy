# INTEL-AMT-PROXY

Transparent http server with automatic login as LSA (Local System Account) for intel-amt development.

# WARNING

Exposing an insecure HTTP Endpoint logged in with the AMT system account!

Be careful about security!

# EXAMPLE

```bash
$ curl -X POST -d '<?xml version="1.0" encoding="utf-8"?>
<Envelope
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing"
	xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
	xmlns="http://www.w3.org/2003/05/soap-envelope">
	<Header>
		<a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action>
		<a:To>/wsman</a:To>
		<w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings</w:ResourceURI>
		<a:MessageID>1</a:MessageID>
		<a:ReplyTo>
			<a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address>
		</a:ReplyTo>
		<w:OperationTimeout>PT60S</w:OperationTimeout>
	</Header>
	<Body></Body>
</Envelope>' http://127.0.0.1:26992/wsman

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope
	xmlns:a="http://www.w3.org/2003/05/soap-envelope"
	xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing"
	xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
	xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust"
	xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd"
	xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<a:Header>
		<b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
		<b:RelatesTo>1</b:RelatesTo>
		<b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</b:Action>
		<b:MessageID>uuid:00000000-8086-8086-8086-00000000006A</b:MessageID>
		<c:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings</c:ResourceURI>
	</a:Header>
	<a:Body>
		<g:AMT_GeneralSettings>
			<g:AMTNetworkEnabled>1</g:AMTNetworkEnabled>
			<g:DDNSPeriodicUpdateInterval>1440</g:DDNSPeriodicUpdateInterval>
			<g:DDNSTTL>900</g:DDNSTTL>
			<g:DDNSUpdateByDHCPServerEnabled>true</g:DDNSUpdateByDHCPServerEnabled>
			<g:DDNSUpdateEnabled>false</g:DDNSUpdateEnabled>
			<g:DHCPSyncRequiresHostname>0</g:DHCPSyncRequiresHostname>
			<g:DHCPv6ConfigurationTimeout>0</g:DHCPv6ConfigurationTimeout>
			<g:DigestRealm>Digest:..........</g:DigestRealm>
			<g:DomainName></g:DomainName>
			<g:ElementName>Intel(r) AMT: General Settings</g:ElementName>
			<g:HostName></g:HostName>
			<g:HostOSFQDN>TEST-PC</g:HostOSFQDN>
			<g:IdleWakeTimeout>1</g:IdleWakeTimeout>
			<g:InstanceID>Intel(r) AMT: General Settings</g:InstanceID>
			<g:NetworkInterfaceEnabled>true</g:NetworkInterfaceEnabled>
			<g:PingResponseEnabled>true</g:PingResponseEnabled>
			<g:PowerSource>0</g:PowerSource>
			<g:PreferredAddressFamily>0</g:PreferredAddressFamily>
			<g:PresenceNotificationInterval>0</g:PresenceNotificationInterval>
			<g:PrivacyLevel>0</g:PrivacyLevel>
			<g:RmcpPingResponseEnabled>false</g:RmcpPingResponseEnabled>
			<g:SharedFQDN>true</g:SharedFQDN>
			<g:ThunderboltDockEnabled>0</g:ThunderboltDockEnabled>
			<g:WsmanOnlyMode>false</g:WsmanOnlyMode>
		</g:AMT_GeneralSettings>
	</a:Body>
</a:Envelope>
```


# License

[GPL-3.0](./LICENSE)
