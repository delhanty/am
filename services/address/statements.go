package address

var queryMap = map[string]string{
	// am.scan_group_addresses related
	"scanGroupAddressesCount": `select count(address_id) as count from am.scan_group_addresses where organization_id=$1 and scan_group_id=$2`,

	// returns
	"scanGroupAddressesAll": `select 
		organization_id, 
		address_id, 
		scan_group_id, 
		host_address,
		ip_address, 
		discovered_timestamp, 
		(select discovered_by from am.scan_address_discovered_by where discovery_id=sga.discovery_id),
		last_scanned_timestamp,
		last_seen_timestamp,
		confidence_score,
		user_confidence_score,
		is_soa,
		is_wildcard_zone,
		is_hosted_service,
		ignored,
		found_from,
		ns_record,
		address_hash
		from am.scan_group_addresses as sga where organization_id=$1 and scan_group_id=$2 and address_id > $3 order by address_id limit $4`,

	"scanGroupAddressesSinceScannedTime": `select 
		organization_id, 
		address_id, 
		scan_group_id, 
		host_address,
		ip_address, 
		discovered_timestamp, 
		(select discovered_by from am.scan_address_discovered_by where discovery_id=sga.discovery_id),
		last_scanned_timestamp,
		last_seen_timestamp,
		confidence_score,
		user_confidence_score,
		is_soa,
		is_wildcard_zone,
		is_hosted_service,
		ignored,
		found_from,
		ns_record,
		address_hash
		from am.scan_group_addresses as sga where organization_id=$1 and scan_group_id=$2 and last_scanned_timestamp > $3 and address_id > $4 order by address_id limit $5`,

	"scanGroupAddressesIgnored": `select 
		organization_id, 
		address_id, 
		scan_group_id, 
		host_address,
		ip_address, 
		discovered_timestamp, 
		(select discovered_by from am.scan_address_discovered_by where discovery_id=sga.discovery_id),
		last_scanned_timestamp,
		last_seen_timestamp,
		confidence_score,
		user_confidence_score,
		is_soa,
		is_wildcard_zone,
		is_hosted_service,
		ignored,
		found_from,
		ns_record,
		address_hash
		from am.scan_group_addresses as sga where organization_id=$1 and scan_group_id=$2 and ignored=$3 and address_id > $4 order by address_id limit $5`,
}

var (
	AddAddressesTempTableKey     = "sga_add_temp"
	AddAddressesTempTableColumns = []string{"address_id", "organization_id", "scan_group_id", "host_address", "ip_address",
		"discovered_timestamp", "discovered_by", "last_scanned_timestamp", "last_seen_timestamp", "confidence_score",
		"user_confidence_score", "is_soa", "is_wildcard_zone", "is_hosted_service", "ignored", "found_from", "ns_record", "address_hash"}
	AddAddressesTempTable = `create temporary table sga_add_temp (
			address_id bigint not null,
			organization_id integer not null,
			scan_group_id integer not null,
			host_address varchar(512),
			ip_address varchar(256),
			discovered_timestamp bigint,
			discovered_by varchar,
			last_scanned_timestamp bigint,
			last_seen_timestamp bigint,
			confidence_score float,
			user_confidence_score float,
			is_soa boolean not null,
			is_wildcard_zone boolean not null,
			is_hosted_service boolean not null,
			ignored boolean not null,
			found_from varchar(128),
			ns_record int,
			address_hash varchar(128)
			check (host_address is not null or ip_address is not null)
		) on commit drop;`

	AddAddressesTempToAddress = `insert into am.scan_group_addresses as sga (
			address_id,
			organization_id, 
			scan_group_id,
			host_address,
			ip_address,
			discovered_timestamp,
			discovery_id,
			last_scanned_timestamp,
			last_seen_timestamp,
			confidence_score,
			user_confidence_score,
			is_soa,
			is_wildcard_zone,
			is_hosted_service,
			ignored,
			found_from,
			ns_record,
			address_hash
		)
		select
			(case when (temp.address_id<>0) then (temp.address_id) else nextval('am.scan_group_addresses_address_id_seq'::regclass) end),
			temp.organization_id, 
			temp.scan_group_id, 
			temp.host_address, 
			temp.ip_address,
			temp.discovered_timestamp, 
			(select discovery_id from am.scan_address_discovered_by where discovered_by=temp.discovered_by),
			temp.last_scanned_timestamp,
			temp.last_seen_timestamp,
			temp.confidence_score,
			temp.user_confidence_score,
			temp.is_soa,
			temp.is_wildcard_zone,
			temp.is_hosted_service,
			temp.ignored,
			temp.found_from,
			temp.ns_record,
			temp.address_hash 
		from sga_add_temp as temp on conflict (scan_group_id, host_address, ip_address) do update set
			last_scanned_timestamp=EXCLUDED.last_scanned_timestamp,
			last_seen_timestamp=EXCLUDED.last_seen_timestamp,
			confidence_score=EXCLUDED.confidence_score,
			user_confidence_score=EXCLUDED.user_confidence_score,
			is_soa=EXCLUDED.is_soa,
			is_wildcard_zone=EXCLUDED.is_wildcard_zone,
			is_hosted_service=EXCLUDED.is_hosted_service,
			ignored=EXCLUDED.ignored,
			found_from=EXCLUDED.found_from,
			ns_record=EXCLUDED.ns_record,
			address_hash=EXCLUDED.address_hash,
			address_id=(case when (sga.address_id<>0) then (sga.address_id) else nextval('am.scan_group_addresses_address_id_seq'::regclass) end);`

	DeleteAddressesTempTableKey     = "sga_del_temp"
	DeleteAddressesTempTableColumns = []string{"address_id"}
	DeleteAddressesTempTable        = `create temporary table sga_del_temp (
			address_id bigint not null
		) on commit drop;`

	DeleteAddressesTempToAddress = `delete from am.scan_group_addresses as sga 
		where organization_id=$1 and scan_group_id=$2 and address_id IN (SELECT address_id FROM sga_del_temp)`
)
