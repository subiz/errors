//go:generate stringer -type=Code -trimprefix=E_
package errors

type Code int

const (
	E_no_error                  Code = 0
	E_unknown                   Code = 1
	E_internal_error            Code = 2
	E_missing_account_id        Code = 10
	E_proto_marshal_error       Code = 11
	E_database_error            Code = 12
	E_access_deny               Code = 13
	E_json_marshal_error        Code = 14
	E_idempotency_input_changed Code = 15
	E_account_not_found         Code = 16
	E_invalid_credential        Code = 17
	E_http_call_error           Code = 18
	E_agent_group_not_found     Code = 20
	E_token_type_is_invalid     Code = 21
	E_jwt_token_is_invalid      Code = 22
	E_agent_not_found           Code = 23
	E_account_fullname_is_empty Code = 24
	E_password_is_empty         Code = 25
	E_password_too_weak         Code = 26
	E_email_is_empty            Code = 27
	E_email_is_invalid          Code = 28
	E_email_is_used             Code = 29
	E_content_type_is_not_json  Code = 30

	E_token_expired                         Code = 31
	E_agent_is_not_active                   Code = 32
	E_wrong_password                        Code = 33
	E_change_own_state                      Code = 34
	E_owner_state_cannot_be_changed         Code = 35
	E_pending_agent_state_cannot_be_changed Code = 36
	E_deleted_agent_state_cannot_be_changed Code = 37
	E_invalid_form_web_type                 Code = 38
	E_invalid_form_mobile_type              Code = 39
	E_invalid_form_email_type               Code = 40
	E_permission_not_enough                 Code = 41
	E_token_is_invalid                      Code = 42
	E_invalid_datetime                      Code = 43

	E_body_parsing_error       Code = 52
	E_query_parsing_error      Code = 53
	E_form_parsing_error       Code = 54
	E_connection_id_is_invalid Code = 55

	E_invalid_filestore_credential  Code = 102
	E_filestore_error               Code = 103
	E_filestore_write_error         Code = 104
	E_filestore_read_error          Code = 105
	E_filestore_acl_error           Code = 106
	E_redis_client_is_uninitialized Code = 150
	E_cannot_connect_to_redis       Code = 151
	E_redis_error                   Code = 152

	E_scrypt_error               Code = 160
	E_randomize_error            Code = 161
	E_scrypt_file_not_found      Code = 162
	E_invalid_mask               Code = 164
	E_form_is_invalid            Code = 165
	E_invalid_access_token       Code = 166
	E_invalid_apikey             Code = 167
	E_invalid_password           Code = 168
	E_whitelist_domain_not_found Code = 170
	E_blacklist_ip_not_found     Code = 171
	E_blocked_user_not_found     Code = 172
	E_user_id_is_invalid         Code = 173
	E_account_id_is_invalid      Code = 174
	E_ip_is_blocked              Code = 175
	E_answer_is_wrong            Code = 176
	E_url_is_not_whitelisted     Code = 177
	E_user_is_banned             Code = 178
	E_domain_is_not_whitelisted  Code = 179

	E_attribute_name_is_empty        Code = 190
	E_attribute_type_is_empty        Code = 191
	E_attribute_key_is_empty         Code = 192
	E_attribyte_type_is_unsupported  Code = 193
	E_attribute_list_is_empty        Code = 194
	E_attribute_key_is_invalid       Code = 195
	E_too_many_attribute             Code = 196
	E_attribute_key_not_found        Code = 197
	E_attribute_list_item_is_invalid Code = 198
	E_attribute_datetime_is_invalid  Code = 199
	E_attribute_boolean_is_invalid   Code = 200
	E_attribute_number_is_invalid    Code = 201
	E_attribute_text_is_invalid      Code = 202

	E_text_too_long Code = 203

	E_automation_not_found       Code = 204
	E_segmentation_not_found     Code = 205
	E_segmentation_id_is_invalid Code = 206
	E_event_id_is_invalid        Code = 207
	E_note_target_id_is_invalid Code = 208
	E_note_id_is_invalid = 209
	E_unknown_segmentation_condition_type = 210

	E_kv_key_not_found                           Code = 1500
	E_subscription_not_found                     Code = 1012
	E_payment_method_not_found                   Code = 1013
	E_missing_primary_payment_method             Code = 1014
	E_payment_method_state_is_failed             Code = 1015
	E_primary_payment_method_is_not_credit_card  Code = 1016
	E_missing_payment_method_id                  Code = 1017
	E_plan_not_found                             Code = 1018
	E_invalid_plan_cannot_buy                    Code = 1019
	E_missing_next_billing_cycle_month           Code = 1020
	E_trial_package_cannot_buy                   Code = 1021
	E_missing_billing_cycle_month                Code = 1022
	E_primary_payment_method_must_be_credit_card Code = 1023
	E_bank_transfer_payment_method_not_found     Code = 1024
	E_invoice_not_found                          Code = 1025
	E_invoice_template_error                     Code = 1026
	E_execute_invoice_template_error             Code = 1027
	E_missing_stripe_token                       Code = 1028
	E_missing_stripe_info                        Code = 1029
	E_stripe_call_error                          Code = 1030
	E_missing_stripe_customer_id                 Code = 1031
	E_invalid_invoice_id                         Code = 1032
	E_invalid_payment_log_id                     Code = 1034
	E_invalid_bill_id                            Code = 1035
	E_unable_to_load_permission                  Code = 1036
	E_auto_charge_is_not_enabled                 Code = 1037
	E_stripe_customer_is_not_valid               Code = 1038
	E_invoice_duedate_empty                      Code = 1039

	E_kafka_rpc_timeout Code = 2150
	E_kafka_error       Code = 2151

	E_s3_call_error                   Code = 1503
	E_pdf_generate_error              Code = 1504
	E_subscription_started_is_invalid Code = 1505
	E_subscription_is_nil             Code = 1506

	E_ticket_list_anchor_is_invalid     Code = 1600
	E_elastic_search_error              Code = 1601
	E_conversation_not_found            Code = 1602
	E_conversation_id_missing           Code = 1603
	E_ticket_id_missing                 Code = 1604
	E_ticket_id_is_invalid              Code = 1605
	E_conversation_id_is_invalid        Code = 1606
	E_conversation_already_had_a_ticket Code = 1607
	E_ticket_not_found                  Code = 1608
	E_webhook_not_found                 Code = 1609
	E_webhook_is_disabled               Code = 1610
	E_conversation_state_is_invalid     Code = 1611
	E_event_is_invalid                  Code = 1612

	E_client_not_found             Code = 3004
	E_missing_redirect_url         Code = 3005
	E_missing_client_name          Code = 3006
	E_corrupted_user_data          Code = 3007
	E_auth_token_expired           Code = 3008
	E_hash_bcrypt_error            Code = 3009
	E_invalid_cipher_size          Code = 3010
	E_invalid_google_auth_response Code = 3011
	E_challenge_not_found          Code = 3054
	E_inactive_agent               Code = 4000
	E_date_format_is_invalid       Code = 4001
	E_pipeline_not_found           Code = 4002
	E_pipeline_is_invalid          Code = 4003
	E_stage_is_invalid             Code = 4004
	E_currency_not_found           Code = 4005
	E_currency_is_invalid          Code = 4006
	E_exchange_rate_not_found      Code = 4007
	E_exchange_rate_is_invalid     Code = 4008
	E_working_day_is_invalid       Code = 4009
	E_holiday_is_invalid           Code = 4010

	E_content_id_missing Code = 4100
)
