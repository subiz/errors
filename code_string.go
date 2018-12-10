// Code generated by "stringer -type=Code -trimprefix=E_"; DO NOT EDIT.

package errors

import "strconv"

const _Code_name = "no_errorunknowninternal_errormissing_account_idproto_marshal_errordatabase_erroraccess_denyjson_marshal_erroridempotency_input_changedaccount_not_foundinvalid_credentialhttp_call_erroragent_group_not_foundtoken_type_is_invalidjwt_token_is_invalidagent_not_foundaccount_fullname_is_emptypassword_is_emptypassword_too_weakemail_is_emptyemail_is_invalidemail_is_usedcontent_type_is_not_jsonmissing_urltoken_expiredagent_is_not_activewrong_passwordchange_own_stateowner_state_cannot_be_changedpending_agent_state_cannot_be_changeddeleted_agent_state_cannot_be_changedinvalid_form_web_typeinvalid_form_mobile_typeinvalid_form_email_typepermission_not_enoughtoken_is_invalidinvalid_datetimebody_parsing_errorquery_parsing_errorform_parsing_errorconnection_id_is_invalidagent_id_is_invalidinvalid_filestore_credentialfilestore_errorfilestore_write_errorfilestore_read_errorfilestore_acl_errorredis_client_is_uninitializedcannot_connect_to_redisredis_errorscrypt_errorrandomize_errorscrypt_file_not_foundinvalid_maskform_is_invalidinvalid_access_tokeninvalid_apikeyinvalid_passwordwhitelist_domain_not_foundblacklist_ip_not_foundblocked_user_not_founduser_id_is_invalidaccount_id_is_invalidip_is_blockedanswer_is_wrongurl_is_not_whitelisteduser_is_banneddomain_is_not_whitelistedattribute_name_is_emptyattribute_type_is_emptyattribute_key_is_emptyattribyte_type_is_unsupportedattribute_list_is_emptyattribute_key_is_invalidtoo_many_attributeattribute_key_not_foundattribute_list_item_is_invalidattribute_datetime_is_invalidattribute_boolean_is_invalidattribute_number_is_invalidattribute_text_is_invalidtext_too_longautomation_not_foundsegmentation_not_foundsegmentation_id_is_invalidevent_id_is_invalidnote_target_id_is_invalidnote_id_is_invalidunknown_segmentation_condition_typekv_key_not_foundsubscription_not_foundpayment_method_not_foundmissing_primary_payment_methodpayment_method_state_is_failedprimary_payment_method_is_not_credit_cardmissing_payment_method_idplan_not_foundinvalid_plan_cannot_buymissing_next_billing_cycle_monthtrial_package_cannot_buymissing_billing_cycle_monthprimary_payment_method_must_be_credit_cardbank_transfer_payment_method_not_foundinvoice_not_foundinvoice_template_errorexecute_invoice_template_errormissing_stripe_tokenmissing_stripe_infostripe_call_errormissing_stripe_customer_idinvalid_invoice_idinvalid_payment_log_idinvalid_bill_idunable_to_load_permissionauto_charge_is_not_enabledstripe_customer_is_not_validinvoice_duedate_emptykafka_rpc_timeoutkafka_errorkafka_rpc_handler_definition_errors3_call_errorpdf_generate_errorsubscription_started_is_invalidsubscription_is_nilticket_list_anchor_is_invalidelastic_search_errorconversation_not_foundinvalid_accepter_must_be_agentconversation_endedconversation_user_is_the_last_one_in_conversationevent_must_be_message_eventconversation_id_missingticket_id_missingticket_id_is_invalidconversation_id_is_invalidconversation_already_had_a_ticketticket_not_foundwebhook_not_foundwebhook_is_disabledinvalid_conversation_stateevent_is_invalidlimit_is_negativeinvalid_integrationmessage_not_founduser_is_not_in_the_conversationduplicated_message_received_errormessage_id_is_invalidintegration_id_is_invalidintegration_state_is_invalidintegration_not_foundundefined_router_condition_keyconversation_closeduser_has_already_in_conversationremover_is_not_agentcaller_is_not_leaverleaver_is_the_last_one_in_conversationempty_messagetoo_large_messageunknown_message_formattoo_many_attachmentstoo_many_fieldstoo_large_attachmenttoo_long_fieldconversation_closer_is_not_in_the_conversationinvalid_channel_idconversation_tag_not_foundclient_not_foundmissing_redirect_urlmissing_client_namecorrupted_user_dataauth_token_expiredhash_bcrypt_errorinvalid_cipher_sizeinvalid_google_auth_responsechallenge_not_foundinactive_agentdate_format_is_invalidpipeline_not_foundpipeline_is_invalidstage_is_invalidcurrency_not_foundcurrency_is_invalidexchange_rate_not_foundexchange_rate_is_invalidworking_day_is_invalidholiday_is_invalidservice_level_agreement_not_foundservice_level_agreement_is_invalidcontent_id_missingfcm_token_not_found"

var _Code_index = [...]uint16{0, 8, 15, 29, 47, 66, 80, 91, 109, 134, 151, 169, 184, 205, 226, 246, 261, 286, 303, 320, 334, 350, 363, 387, 398, 411, 430, 444, 460, 489, 526, 563, 584, 608, 631, 652, 668, 684, 702, 721, 739, 763, 782, 810, 825, 846, 866, 885, 914, 937, 948, 960, 975, 996, 1008, 1023, 1043, 1057, 1073, 1099, 1121, 1143, 1161, 1182, 1195, 1210, 1232, 1246, 1271, 1294, 1317, 1339, 1368, 1391, 1415, 1433, 1456, 1486, 1515, 1543, 1570, 1595, 1608, 1628, 1650, 1676, 1695, 1720, 1738, 1773, 1789, 1811, 1835, 1865, 1895, 1936, 1961, 1975, 1998, 2030, 2054, 2081, 2123, 2161, 2178, 2200, 2230, 2250, 2269, 2286, 2312, 2330, 2352, 2367, 2392, 2418, 2446, 2467, 2484, 2495, 2529, 2542, 2560, 2591, 2610, 2639, 2659, 2681, 2711, 2729, 2778, 2805, 2828, 2845, 2865, 2891, 2924, 2940, 2957, 2976, 3002, 3018, 3035, 3054, 3071, 3102, 3135, 3156, 3181, 3209, 3230, 3260, 3279, 3311, 3331, 3351, 3389, 3402, 3419, 3441, 3461, 3476, 3496, 3510, 3556, 3574, 3600, 3616, 3636, 3655, 3674, 3692, 3709, 3728, 3756, 3775, 3789, 3811, 3829, 3848, 3864, 3882, 3901, 3924, 3948, 3970, 3988, 4021, 4055, 4073, 4092}

func (i Code) String() string {
	if i < 0 || i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}
