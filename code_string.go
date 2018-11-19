// Code generated by "stringer -type=Code -trimprefix=E_"; DO NOT EDIT.

package errors

import "strconv"

const _Code_name = "no_errorunknowninternal_errormissing_account_idproto_marshal_errordatabase_erroraccess_denyjson_marshal_erroridempotency_input_changedaccount_not_foundinvalid_credentialhttp_call_erroragent_group_not_foundtoken_type_is_invalidjwt_token_is_invalidagent_not_foundaccount_fullname_is_emptypassword_is_emptypassword_too_weakemail_is_emptyemail_is_invalidemail_is_usedcontent_type_is_not_jsontoken_expiredagent_is_not_activewrong_passwordchange_own_stateowner_state_cannot_be_changedpending_agent_state_cannot_be_changeddeleted_agent_state_cannot_be_changedinvalid_form_web_typeinvalid_form_mobile_typeinvalid_form_email_typepermission_not_enoughbody_parsing_errorquery_parsing_errorform_parsing_errorconnection_id_is_invalidinvalid_filestore_credentialfilestore_errorfilestore_write_errorfilestore_read_errorfilestore_acl_errorredis_client_is_uninitializedcannot_connect_to_redisredis_errorscrypt_errorrandomize_errorscrypt_file_not_foundinvalid_maskform_is_invalidinvalid_access_tokeninvalid_apikeyinvalid_passwordwhitelist_domain_not_foundblacklist_ip_not_foundblocked_user_not_founduser_id_is_invalidaccount_id_is_invalidip_is_blockedanswer_is_wrongsubscription_not_foundpayment_method_not_foundmissing_primary_payment_methodpayment_method_state_is_failedprimary_payment_method_is_not_credit_cardmissing_payment_method_idplan_not_foundinvalid_plan_cannot_buymissing_next_billing_cycle_monthtrial_package_cannot_buymissing_billing_cycle_monthprimary_payment_method_must_be_credit_cardbank_transfer_payment_method_not_foundinvoice_not_foundinvoice_template_errorexecute_invoice_template_errormissing_stripe_tokenmissing_stripe_infostripe_call_errormissing_stripe_customer_idinvalid_invoice_idinvalid_payment_log_idinvalid_bill_idunable_to_load_permissionauto_charge_is_not_enabledstripe_customer_is_not_validinvoice_duedate_emptykv_key_not_founds3_call_errorpdf_generate_errorsubscription_started_is_invalidsubscription_is_nilticket_list_anchor_is_invalidelastic_search_errorconversation_not_foundconversation_id_missingticket_id_missingticket_id_is_invalidconversation_id_is_invalidconversation_already_had_a_ticketticket_not_foundwebhook_not_foundwebhook_is_disabledconversation_state_is_invalidevent_is_invalidkafka_rpc_timeoutclient_not_foundmissing_redirect_urlmissing_client_namecorrupted_user_dataauth_token_expiredhash_bcrypt_errorinvalid_cipher_sizeinvalid_google_auth_responsechallenge_not_foundinactive_agentdate_format_is_invalidpipeline_not_foundpipeline_is_invalidstage_is_invalidcurrency_not_foundcurrency_is_invalidexchange_rate_not_foundexchange_rate_is_invalidworking_day_is_invalidholiday_is_invalidcontent_id_missing"

var _Code_map = map[Code]string{
	0:    _Code_name[0:8],
	1:    _Code_name[8:15],
	2:    _Code_name[15:29],
	10:   _Code_name[29:47],
	11:   _Code_name[47:66],
	12:   _Code_name[66:80],
	13:   _Code_name[80:91],
	14:   _Code_name[91:109],
	15:   _Code_name[109:134],
	16:   _Code_name[134:151],
	17:   _Code_name[151:169],
	18:   _Code_name[169:184],
	20:   _Code_name[184:205],
	21:   _Code_name[205:226],
	22:   _Code_name[226:246],
	23:   _Code_name[246:261],
	24:   _Code_name[261:286],
	25:   _Code_name[286:303],
	26:   _Code_name[303:320],
	27:   _Code_name[320:334],
	28:   _Code_name[334:350],
	29:   _Code_name[350:363],
	30:   _Code_name[363:387],
	31:   _Code_name[387:400],
	32:   _Code_name[400:419],
	33:   _Code_name[419:433],
	34:   _Code_name[433:449],
	35:   _Code_name[449:478],
	36:   _Code_name[478:515],
	37:   _Code_name[515:552],
	38:   _Code_name[552:573],
	39:   _Code_name[573:597],
	40:   _Code_name[597:620],
	41:   _Code_name[620:641],
	52:   _Code_name[641:659],
	53:   _Code_name[659:678],
	54:   _Code_name[678:696],
	55:   _Code_name[696:720],
	102:  _Code_name[720:748],
	103:  _Code_name[748:763],
	104:  _Code_name[763:784],
	105:  _Code_name[784:804],
	106:  _Code_name[804:823],
	150:  _Code_name[823:852],
	151:  _Code_name[852:875],
	152:  _Code_name[875:886],
	160:  _Code_name[886:898],
	161:  _Code_name[898:913],
	162:  _Code_name[913:934],
	164:  _Code_name[934:946],
	165:  _Code_name[946:961],
	166:  _Code_name[961:981],
	167:  _Code_name[981:995],
	168:  _Code_name[995:1011],
	170:  _Code_name[1011:1037],
	171:  _Code_name[1037:1059],
	172:  _Code_name[1059:1081],
	173:  _Code_name[1081:1099],
	174:  _Code_name[1099:1120],
	175:  _Code_name[1120:1133],
	176:  _Code_name[1133:1148],
	1012: _Code_name[1148:1170],
	1013: _Code_name[1170:1194],
	1014: _Code_name[1194:1224],
	1015: _Code_name[1224:1254],
	1016: _Code_name[1254:1295],
	1017: _Code_name[1295:1320],
	1018: _Code_name[1320:1334],
	1019: _Code_name[1334:1357],
	1020: _Code_name[1357:1389],
	1021: _Code_name[1389:1413],
	1022: _Code_name[1413:1440],
	1023: _Code_name[1440:1482],
	1024: _Code_name[1482:1520],
	1025: _Code_name[1520:1537],
	1026: _Code_name[1537:1559],
	1027: _Code_name[1559:1589],
	1028: _Code_name[1589:1609],
	1029: _Code_name[1609:1628],
	1030: _Code_name[1628:1645],
	1031: _Code_name[1645:1671],
	1032: _Code_name[1671:1689],
	1034: _Code_name[1689:1711],
	1035: _Code_name[1711:1726],
	1036: _Code_name[1726:1751],
	1037: _Code_name[1751:1777],
	1038: _Code_name[1777:1805],
	1039: _Code_name[1805:1826],
	1500: _Code_name[1826:1842],
	1503: _Code_name[1842:1855],
	1504: _Code_name[1855:1873],
	1505: _Code_name[1873:1904],
	1506: _Code_name[1904:1923],
	1600: _Code_name[1923:1952],
	1601: _Code_name[1952:1972],
	1602: _Code_name[1972:1994],
	1603: _Code_name[1994:2017],
	1604: _Code_name[2017:2034],
	1605: _Code_name[2034:2054],
	1606: _Code_name[2054:2080],
	1607: _Code_name[2080:2113],
	1608: _Code_name[2113:2129],
	1609: _Code_name[2129:2146],
	1610: _Code_name[2146:2165],
	1611: _Code_name[2165:2194],
	1612: _Code_name[2194:2210],
	2150: _Code_name[2210:2227],
	3004: _Code_name[2227:2243],
	3005: _Code_name[2243:2263],
	3006: _Code_name[2263:2282],
	3007: _Code_name[2282:2301],
	3008: _Code_name[2301:2319],
	3009: _Code_name[2319:2336],
	3010: _Code_name[2336:2355],
	3011: _Code_name[2355:2383],
	3054: _Code_name[2383:2402],
	4000: _Code_name[2402:2416],
	4001: _Code_name[2416:2438],
	4002: _Code_name[2438:2456],
	4003: _Code_name[2456:2475],
	4004: _Code_name[2475:2491],
	4005: _Code_name[2491:2509],
	4006: _Code_name[2509:2528],
	4007: _Code_name[2528:2551],
	4008: _Code_name[2551:2575],
	4009: _Code_name[2575:2597],
	4010: _Code_name[2597:2615],
	4100: _Code_name[2615:2633],
}

func (i Code) String() string {
	if str, ok := _Code_map[i]; ok {
		return str
	}
	return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
}
