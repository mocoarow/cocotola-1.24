load("@gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "cat_dario_mergo",
        importpath = "dario.cat/mergo",
        sum = "h1:85+piFYR1tMbRrLcDwR18y4UKJ3aH1Tbzi24VRW1TK8=",
        version = "v1.0.2",
    )
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:qTakTkI6ni6LFD5sBwwsdSO+AQqbSIxOauHTTQKZ/7o=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_99designs_go_keychain",
        importpath = "github.com/99designs/go-keychain",
        sum = "h1:/vQbFIOMbk2FiG/kXiLl8BRyzTWDw7gX/Hz7Dd5eDMs=",
        version = "v0.0.0-20191008050251-8e49817e8af4",
    )
    go_repository(
        name = "com_github_99designs_keyring",
        importpath = "github.com/99designs/keyring",
        sum = "h1:tYLp1ULvO7i3fI5vE21ReQuj99QFSs7lGm0xWyJo87o=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_ajg_form",
        importpath = "github.com/ajg/form",
        sum = "h1:t9c7v8JUKu/XxOGBU0yjNpaMloxGEJhUkqFRq0ibGeU=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_ajstarks_deck",
        importpath = "github.com/ajstarks/deck",
        sum = "h1:7kQgkwGRoLzC9K0oyXdJo7nve/bynv/KwUsxbiTlzAM=",
        version = "v0.0.0-20200831202436-30c9fc6549a9",
    )
    go_repository(
        name = "com_github_ajstarks_deck_generate",
        importpath = "github.com/ajstarks/deck/generate",
        sum = "h1:iXUgAaqDcIUGbRoy2TdeofRG/j1zpGRSEmNK05T+bi8=",
        version = "v0.0.0-20210309230005-c3f852c02e19",
    )
    go_repository(
        name = "com_github_ajstarks_svgo",
        importpath = "github.com/ajstarks/svgo",
        sum = "h1:slYM766cy2nI3BwyRiyQj/Ud48djTMtMebDqepE95rw=",
        version = "v0.0.0-20211024235047-1546f124cd8b",
    )
    go_repository(
        name = "com_github_alecaivazis_survey_v2",
        importpath = "github.com/AlecAivazis/survey/v2",
        sum = "h1:6I/u8FvytdGsgonrYsVn2t8t4QiRnh6QSTqkkhIiSjQ=",
        version = "v2.3.7",
    )
    go_repository(
        name = "com_github_alecthomas_chroma_v2",
        importpath = "github.com/alecthomas/chroma/v2",
        sum = "h1:R3+wzpnUArGcQz7fCETQBzO5n9IMNi13iIs46aU4V9E=",
        version = "v2.14.0",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin_v2",
        importpath = "github.com/alecthomas/kingpin/v2",
        sum = "h1:f48lwail6p8zpO1bC4TxtqACaGqHYA22qkHjHpqDjYY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        importpath = "github.com/alecthomas/units",
        sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
        version = "v0.0.0-20211218093645-b94a6e3cc137",
    )
    go_repository(
        name = "com_github_andybalholm_brotli",
        importpath = "github.com/andybalholm/brotli",
        sum = "h1:V7DdXeJtZscaqfNuAdSRuRFzuiKlHSC/Zh3zl9qY3JY=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_anmitsu_go_shlex",
        importpath = "github.com/anmitsu/go-shlex",
        sum = "h1:9AeTilPcZAjCFIImctFaOjnTIavg87rW78vTPkQqLI8=",
        version = "v0.0.0-20200514113438-38f4b401e2be",
    )
    go_repository(
        name = "com_github_antihax_optional",
        importpath = "github.com/antihax/optional",
        sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v10",
        importpath = "github.com/apache/arrow/go/v10",
        sum = "h1:n9dERvixoC/1JjDmBcs9FPaEryoANa2sCgVFo6ez9cI=",
        version = "v10.0.1",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v11",
        importpath = "github.com/apache/arrow/go/v11",
        sum = "h1:hqauxvFQxww+0mEU/2XHG6LT7eZternCZq+A5Yly2uM=",
        version = "v11.0.0",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v15",
        importpath = "github.com/apache/arrow/go/v15",
        sum = "h1:60IliRbiyTWCWjERBCkO1W4Qun9svcYoZrSLcyOsMLE=",
        version = "v15.0.2",
    )
    go_repository(
        name = "com_github_apache_thrift",
        importpath = "github.com/apache/thrift",
        sum = "h1:qEy6UW60iVOlUy+b9ZR0d5WzUWYGOo4HfopoyBaNmoY=",
        version = "v0.16.0",
    )
    go_repository(
        name = "com_github_araddon_dateparse",
        importpath = "github.com/araddon/dateparse",
        sum = "h1:FxWPpzIjnTlhPwqqXc4/vE0f7GvRjuAsbW+HOIe8KnA=",
        version = "v0.0.0-20210429162001-6b43995a97de",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:yNldzF5kzLBRvKlKz1S0bkvc2+04R1kt13KfBWQBfFA=",
        version = "v1.49.6",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:mJoei2CxPutQVxaATCzDUjcZEjVRdpsiiXi2o38yqWM=",
        version = "v1.36.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_aws_protocol_eventstream",
        importpath = "github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream",
        sum = "h1:tcFliCWne+zOuUfKNRn8JdFBuWPDuISDH08wD2ULkhk=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_config",
        importpath = "github.com/aws/aws-sdk-go-v2/config",
        sum = "h1:Y/2a+jLPrPbHpFkpAAYkVEtJmxORlXoo5k2g1fa2sUo=",
        version = "v1.29.12",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_credentials",
        importpath = "github.com/aws/aws-sdk-go-v2/credentials",
        sum = "h1:q+nV2yYegofO/SUXruT+pn4KxkxmaQ++1B/QedcKBFM=",
        version = "v1.17.65",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
        sum = "h1:x793wxmUWVDhshP8WW2mlnXuFrO4cOd3HLBroh1paFw=",
        version = "v1.16.30",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_s3_manager",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/s3/manager",
        sum = "h1:fAoVmNGhir6BR+RU0/EI+6+D7abM+MCwWf8v4ip5jNI=",
        version = "v1.11.33",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
        sum = "h1:ZK5jHhnrioRkUNOc+hOgQKlUL5JeC3S6JgLxtQ+Rm0Q=",
        version = "v1.3.34",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
        sum = "h1:SZwFm17ZUNNg5Np0ioo/gq8Mn6u9w19Mri8DnJ15Jf0=",
        version = "v2.6.34",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_ini",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/ini",
        sum = "h1:bIqFDwgGXXN1Kpp99pDOdKMTTb5d2KyU5X/BZxjOkRo=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_v4a",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/v4a",
        sum = "h1:ZSIPAkAsCCjYrhqfw2+lNzWDzxzHXEckFkTePL5RSWQ=",
        version = "v1.0.14",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_accept_encoding",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding",
        sum = "h1:eAh2A4b5IzM/lum78bZ590jy36+d/aFLgKF/4Vd1xPE=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_checksum",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/checksum",
        sum = "h1:BBYoNQt2kUZUUK4bIPsKrCcjVPUMNsgQpNAwhznK/zo=",
        version = "v1.1.18",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
        sum = "h1:dM9/92u2F1JbDaGooxTq18wmmFzbJRfXfVfy96/1CXM=",
        version = "v1.12.15",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_s3shared",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/s3shared",
        sum = "h1:HfVVR1vItaG6le+Bpw6P4midjBDMKnjMyZnw9MXYUcE=",
        version = "v1.13.17",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_s3",
        importpath = "github.com/aws/aws-sdk-go-v2/service/s3",
        sum = "h1:3/gm/JTX9bX8CpzTgIlrtYpB3EVBDxyg/GY/QdcIEZw=",
        version = "v1.27.11",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sso",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
        sum = "h1:pdgODsAhGo4dvzC3JAG5Ce0PX8kWXrTZGx+jxADD+5E=",
        version = "v1.25.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_ssooidc",
        importpath = "github.com/aws/aws-sdk-go-v2/service/ssooidc",
        sum = "h1:90uX0veLKcdHVfvxhkWUQSCi5VabtwMLFutYiRke4oo=",
        version = "v1.30.0",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sts",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
        sum = "h1:PZV5W8yk4OtH1JAuhV2PXwwO9v5G5Aoj+eMCn4T+1Kc=",
        version = "v1.33.17",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:Z//5NuZCSW6R4PhQ93hShNbyBbn8BWCmCVCt+Q8Io5k=",
        version = "v1.22.3",
    )
    go_repository(
        name = "com_github_aybabtme_uniplot",
        importpath = "github.com/aybabtme/uniplot",
        sum = "h1:dSeuFcs4WAJJnswS8vXy7YY1+fdlbVPuEVmDAfqvFOQ=",
        version = "v0.0.0-20151203143629-039c559e5e7e",
    )
    go_repository(
        name = "com_github_aymanbagabas_go_osc52_v2",
        importpath = "github.com/aymanbagabas/go-osc52/v2",
        sum = "h1:HwpRHbFMcZLEVr42D4p7XBqjyuxQH5SMiErDT4WkJ2k=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_aymerick_douceur",
        importpath = "github.com/aymerick/douceur",
        sum = "h1:Mv+mAeH1Q+n9Fr+oyamOlAkUNPWPlA8PPGR0QAaYuPk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_azcore",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/azcore",
        sum = "h1:Gt0j3wceWMwPmiazCa8MzMA0MfhmPIz0Qp0FJ6qcM0U=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_azidentity",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/azidentity",
        sum = "h1:B+blDbyVIG3WaikNxPnhPiJ1MThR03b3vKGtER95TP4=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_internal",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/internal",
        sum = "h1:FPKJS1T+clwv+OLGt13a8UjqeRuh0O4SJ3lUriThc+4=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_security_keyvault_azkeys",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys",
        sum = "h1:Wgf5rZba3YZqeTNJPtvqZoBu1sBN/L4sry+u2U3Y75w=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_security_keyvault_internal",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal",
        sum = "h1:bFWuoEKg+gImo7pvkiQEFAc8ocibADgXeiLAxWhWmkI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_storage_azblob",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob",
        sum = "h1:u/LLAOFgsMv7HmNL4Qufg58y+qElGOt5qv0z1mURkRY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        importpath = "github.com/Azure/go-ansiterm",
        sum = "h1:udKWzYgxTojEKWjV8V+WSxDXJ4NFATAsZjh8iIbsQIg=",
        version = "v0.0.0-20250102033503-faa5f7b0171c",
    )
    go_repository(
        name = "com_github_azure_go_autorest",
        importpath = "github.com/Azure/go-autorest",
        sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
        version = "v14.2.0+incompatible",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_adal",
        importpath = "github.com/Azure/go-autorest/autorest/adal",
        sum = "h1:P8An8Z9rH1ldbOLdFpxYorgOt2sywL9V24dAwWHPuGc=",
        version = "v0.9.16",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_date",
        importpath = "github.com/Azure/go-autorest/autorest/date",
        sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_azure_go_autorest_logger",
        importpath = "github.com/Azure/go-autorest/logger",
        sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_tracing",
        importpath = "github.com/Azure/go-autorest/tracing",
        sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_azure_go_ntlmssp",
        importpath = "github.com/Azure/go-ntlmssp",
        sum = "h1:mFRzDkZVAjdal+s7s0MwaRv9igoPqLRdzOLzw/8Xvq8=",
        version = "v0.0.0-20221128193559-754e69321358",
    )
    go_repository(
        name = "com_github_azuread_microsoft_authentication_library_for_go",
        importpath = "github.com/AzureAD/microsoft-authentication-library-for-go",
        sum = "h1:oygO0locgZJe7PpYPXT5A29ZkwJaPqcva7BVeemZOZs=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_bahlo_generic_list_go",
        importpath = "github.com/bahlo/generic-list-go",
        sum = "h1:5sz/EEAK+ls5wF+NeqDpk5+iNdMDXrh3z3nPnH1Wvgk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_benbjohnson_clock",
        importpath = "github.com/benbjohnson/clock",
        sum = "h1:VvXlSJBzZpA/zum6Sj74hxwYI2DIxRWuNIoXAzHZz5o=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bmatcuk_doublestar_v4",
        importpath = "github.com/bmatcuk/doublestar/v4",
        sum = "h1:X8jg9rRZmJd4yRy7ZeNDRnM+T3ZfHv15JiBJ/avrEXE=",
        version = "v4.9.1",
    )
    go_repository(
        name = "com_github_bmizerany_assert",
        importpath = "github.com/bmizerany/assert",
        sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
        version = "v0.0.0-20160611221934-b7ed37b82869",
    )
    go_repository(
        name = "com_github_boombuler_barcode",
        importpath = "github.com/boombuler/barcode",
        sum = "h1:NDBbPmhS+EqABEs5Kg3n/5ZNjy73Pz7SIV+KCeqyXcs=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bradleyfalzon_ghinstallation_v2",
        importpath = "github.com/bradleyfalzon/ghinstallation/v2",
        sum = "h1:B91r9bHtXp/+XRgS5aZm6ZzTdz3ahgJYmkt4xZkgDz8=",
        version = "v2.16.0",
    )
    go_repository(
        name = "com_github_brianvoe_gofakeit_v6",
        importpath = "github.com/brianvoe/gofakeit/v6",
        sum = "h1:Xib46XXuQfmlLS2EXRuJpqcw8St6qSZz75OUo0tgAW4=",
        version = "v6.28.0",
    )
    go_repository(
        name = "com_github_bufbuild_protocompile",
        importpath = "github.com/bufbuild/protocompile",
        sum = "h1:iA73zAf/fyljNjQKwYzUHD6AD4R8KMasmwa/FBatYVw=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_github_buger_jsonparser",
        importpath = "github.com/buger/jsonparser",
        sum = "h1:2PnMjfWD7wBILjqQbt530v576A/cAbQvEW9gGIpYMUs=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_buildkite_interpolate",
        importpath = "github.com/buildkite/interpolate",
        sum = "h1:v2Ji3voik69UZlbfoqzx+qfcsOKLA61nHdU79VV+tPU=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_burntsushi_xgb",
        importpath = "github.com/BurntSushi/xgb",
        sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
        version = "v0.0.0-20160522181843-27f122750802",
    )
    go_repository(
        name = "com_github_bytedance_sonic",
        importpath = "github.com/bytedance/sonic",
        sum = "h1:/OfKt8HFw0kh2rj8N0F6C/qPGRESq0BbaNZgcNXXzQQ=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_bytedance_sonic_loader",
        importpath = "github.com/bytedance/sonic/loader",
        sum = "h1:dskwH8edlzNMctoruo8FPTJDF3vLtDT0sXZwvZJyqeA=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_campoy_embedmd",
        importpath = "github.com/campoy/embedmd",
        sum = "h1:V4kI2qTJJLf4J29RzI/MAt2c3Bl4dQSYPuflzwFH2hY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_casbin_casbin_v2",
        importpath = "github.com/casbin/casbin/v2",
        sum = "h1:Mo9R/EKZk9aoagFs0OmuCmBYjWJfvbWJiX4aenIJOKY=",
        version = "v2.120.0",
    )
    go_repository(
        name = "com_github_casbin_gorm_adapter_v3",
        importpath = "github.com/casbin/gorm-adapter/v3",
        sum = "h1:CeW9R9SeWTnf7JZQ4zlOWBXKcspT1VoSf4ojeFX2IbM=",
        version = "v3.36.0",
    )
    go_repository(
        name = "com_github_casbin_govaluate",
        importpath = "github.com/casbin/govaluate",
        sum = "h1:XB53bSw+gaQ7tjTlFJsuTThPCQBxyUeQZ3drsKiicEY=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:MyRJ/UdXutAwSAT+s3wNd7MfTIcy71VQueUuFK343L8=",
        version = "v4.3.0",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v5",
        importpath = "github.com/cenkalti/backoff/v5",
        sum = "h1:ZN+IMa753KfX5hd8vVaMixjnqRZ3y8CuJKRKj1xcsSM=",
        version = "v5.0.3",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash",
        importpath = "github.com/cespare/xxhash",
        sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:UL815xU9SqsFlibzuggzjXhog7bL6oX9BbNZnL2UFvs=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_charmbracelet_bubbles",
        importpath = "github.com/charmbracelet/bubbles",
        sum = "h1:9TdC97SdRVg/1aaXNVWfFH3nnLAwOXr8Fn6u6mfQdFs=",
        version = "v0.21.0",
    )
    go_repository(
        name = "com_github_charmbracelet_bubbletea",
        importpath = "github.com/charmbracelet/bubbletea",
        sum = "h1:VkHIxPJQeDt0aFJIsVxw8BQdh/F/L2KKZGsK6et5taU=",
        version = "v1.3.6",
    )
    go_repository(
        name = "com_github_charmbracelet_colorprofile",
        importpath = "github.com/charmbracelet/colorprofile",
        sum = "h1:4pZI35227imm7yK2bGPcfpFEmuY1gc2YSTShr4iJBfs=",
        version = "v0.2.3-0.20250311203215-f60798e515dc",
    )
    go_repository(
        name = "com_github_charmbracelet_glamour",
        importpath = "github.com/charmbracelet/glamour",
        sum = "h1:hx6E25SvI2WiZdt/gxINcYBnHD7PE2Vr9auqwg5B05g=",
        version = "v0.9.2-0.20250319212134-549f544650e3",
    )
    go_repository(
        name = "com_github_charmbracelet_lipgloss",
        importpath = "github.com/charmbracelet/lipgloss",
        sum = "h1:nFRtCfZu/zkltd2lsLUPlVNv3ej/Atod9hcdbRZtlys=",
        version = "v1.1.1-0.20250319133953-166f707985bc",
    )
    go_repository(
        name = "com_github_charmbracelet_x_ansi",
        importpath = "github.com/charmbracelet/x/ansi",
        sum = "h1:BXt5DHS/MKF+LjuK4huWrC6NCvHtexww7dMayh6GXd0=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_github_charmbracelet_x_cellbuf",
        importpath = "github.com/charmbracelet/x/cellbuf",
        sum = "h1:/KBBKHuVRbq1lYx5BzEHBAFBP8VcQzJejZ/IA3iR28k=",
        version = "v0.0.13",
    )
    go_repository(
        name = "com_github_charmbracelet_x_term",
        importpath = "github.com/charmbracelet/x/term",
        sum = "h1:AQeHeLZ1OqSXhrAWpYUtZyX1T3zVxfpZuEQMIQaGIAQ=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_cheggaaa_pb",
        importpath = "github.com/cheggaaa/pb",
        sum = "h1:FckUN5ngEk2LpvuG0fw1GEFx6LtyY2pWI/Z2QgCnEYo=",
        version = "v1.0.29",
    )
    go_repository(
        name = "com_github_chromedp_cdproto",
        importpath = "github.com/chromedp/cdproto",
        sum = "h1:UQ4AU+BGti3Sy/aLU8KVseYKNALcX9UXY6DfpwQ6J8E=",
        version = "v0.0.0-20250724212937-08a3db8b4327",
    )
    go_repository(
        name = "com_github_chromedp_chromedp",
        importpath = "github.com/chromedp/chromedp",
        sum = "h1:0uAbnxewy/Q+Bg7oafVePE/6EXEho9hnaC38f+TTENg=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_github_chromedp_sysutil",
        importpath = "github.com/chromedp/sysutil",
        sum = "h1:PUFNv5EcprjqXZD9nJb9b/c9ibAbxiYo4exNWZyipwM=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_chzyer_logex",
        importpath = "github.com/chzyer/logex",
        sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
        version = "v1.1.10",
    )
    go_repository(
        name = "com_github_chzyer_readline",
        importpath = "github.com/chzyer/readline",
        sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
        version = "v0.0.0-20180603132655-2972be24d48e",
    )
    go_repository(
        name = "com_github_chzyer_test",
        importpath = "github.com/chzyer/test",
        sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
        version = "v0.0.0-20180213035817-a1ea475d72b1",
    )
    go_repository(
        name = "com_github_cli_browser",
        importpath = "github.com/cli/browser",
        sum = "h1:LejqCrpWr+1pRqmEPDGnTZOjsMe7sehifLynZJuqJpo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_cli_go_gh_v2",
        importpath = "github.com/cli/go-gh/v2",
        sum = "h1:SVt1/afj5FRAythyMV3WJKaUfDNsxXTIe7arZbwTWKA=",
        version = "v2.12.1",
    )
    go_repository(
        name = "com_github_cli_safeexec",
        importpath = "github.com/cli/safeexec",
        sum = "h1:e/C79PbXF4yYTN/wauC4tviMxEV13BwljGj0N9j+N00=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_cli_shurcool_graphql",
        importpath = "github.com/cli/shurcooL-graphql",
        sum = "h1:6MogPnQJLjKkaXPyGqPRXOI2qCsQdqNfUY1QSJu2GuY=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_clickhouse_clickhouse_go",
        importpath = "github.com/ClickHouse/clickhouse-go",
        sum = "h1:iAFMa2UrQdR5bHJ2/yaSLffZkxpcOYQMCUuKeNXGdqc=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cloudflare_golz4",
        importpath = "github.com/cloudflare/golz4",
        sum = "h1:F1EaeKL/ta07PY/k9Os/UFtwERei2/XzGemhpGnBKNg=",
        version = "v0.0.0-20150217214814-ef862a3cdc58",
    )
    go_repository(
        name = "com_github_cloudwego_base64x",
        importpath = "github.com/cloudwego/base64x",
        sum = "h1:t11wG9AECkCDk5fMSoxmufanudBtJ+/HemLstXDLI2M=",
        version = "v0.1.6",
    )
    go_repository(
        name = "com_github_cloudwego_iasm",
        importpath = "github.com/cloudwego/iasm",
        sum = "h1:1KNIy1I1H9hNNFEEH3DVnI4UujN+1zjpuk6gwHLTssg=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:aQ3y1lwWyqYPiWZThqv1aFbZMiM9vblcSArJRf2Irls=",
        version = "v0.0.0-20250501225837-2ac532fd4443",
    )
    go_repository(
        name = "com_github_cockroachdb_cockroach_go_v2",
        importpath = "github.com/cockroachdb/cockroach-go/v2",
        sum = "h1:3XzfSMuUT0wBe1a3o5C0eOTcArhmmFAg2Jzh/7hhKqo=",
        version = "v2.1.1",
    )
    go_repository(
        name = "com_github_containerd_continuity",
        importpath = "github.com/containerd/continuity",
        sum = "h1:ZRoN1sXq9u7V6QoHMcVWGhOwDFqZ4B9i5H6un1Wh0x4=",
        version = "v0.4.5",
    )
    go_repository(
        name = "com_github_containerd_errdefs",
        importpath = "github.com/containerd/errdefs",
        sum = "h1:tg5yIfIlQIrxYtu9ajqY42W3lpS19XqdxRQeEwYG8PI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_errdefs_pkg",
        importpath = "github.com/containerd/errdefs/pkg",
        sum = "h1:9IKJ06FvyNlexW690DXuQNx2KA2cUJXx151Xdx3ZPPE=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_containerd_typeurl_v2",
        importpath = "github.com/containerd/typeurl/v2",
        sum = "h1:6NBDbQzr7I5LHgp34xAXYF5DOTQDn05X58lsPEmzLso=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:yi21YpKnrx1gt5R+la8n5WgS0kCrsPp33dmEyHReZr4=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:RrqgGjYQKalulkV8NGVIfkXQf6YYmOyiJKk8iXXhfZs=",
        version = "v22.5.0",
    )
    go_repository(
        name = "com_github_creack_pty",
        importpath = "github.com/creack/pty",
        sum = "h1:n56/Zwd5o6whRC5PMGretI4IdRLlmBXYNjScPaBgsbY=",
        version = "v1.1.18",
    )
    go_repository(
        name = "com_github_cznic_mathutil",
        importpath = "github.com/cznic/mathutil",
        sum = "h1:XNT/Zf5l++1Pyg08/HV04ppB0gKxAqtZQBRYiYrUuYk=",
        version = "v0.0.0-20180504122225-ca4c9f2c1369",
    )
    go_repository(
        name = "com_github_danieljoos_wincred",
        importpath = "github.com/danieljoos/wincred",
        sum = "h1:QLdCxFs1/Yl4zduvBdcHB8goaYk9RARS2SgLLRuAyr0=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:U9qPSI2PIWSS1VwoXQT9A3Wy9MM3WgvqSxFWenqJduM=",
        version = "v1.1.2-0.20180830191138-d8f796af33cc",
    )
    go_repository(
        name = "com_github_dchest_bcrypt_pbkdf",
        importpath = "github.com/dchest/bcrypt_pbkdf",
        sum = "h1:saTgr5tMLFnmy/yg3qDTft4rE5DY2uJ/cCxCe3q0XTU=",
        version = "v0.0.0-20150205184540-83f37f9c154a",
    )
    go_repository(
        name = "com_github_decred_dcrd_dcrec_secp256k1_v4",
        importpath = "github.com/decred/dcrd/dcrec/secp256k1/v4",
        sum = "h1:HbphB4TFFXpv7MNrT52FGrrgVXF1owhMVTHFZIlnvd4=",
        version = "v4.1.0",
    )
    go_repository(
        name = "com_github_dhui_dktest",
        importpath = "github.com/dhui/dktest",
        sum = "h1:uUfYBIVREmj/Rw6MvgmqNAYzTiKOHJak+enB5Di73MM=",
        version = "v0.4.5",
    )
    go_repository(
        name = "com_github_distribution_reference",
        importpath = "github.com/distribution/reference",
        sum = "h1:0IXCQ5g4/QMHHkarYzh5l+u8T3t73zM5QvfrDyIgxBk=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_dlclark_regexp2",
        importpath = "github.com/dlclark/regexp2",
        sum = "h1:Q/sSnsKerHeCkc/jSTNq1oCm7KiVgUMZRDUoRu0JQZQ=",
        version = "v1.11.5",
    )
    go_repository(
        name = "com_github_dnaeon_go_vcr",
        importpath = "github.com/dnaeon/go-vcr",
        sum = "h1:zHCHvJYTMh1N7xnV7zf1m1GPBF9Ad0Jk/whtQ1663qI=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_docker_cli",
        importpath = "github.com/docker/cli",
        sum = "h1:fp9ZHAr1WWPGdIWBM1b3zLtgCF+83gRdVMTJsUeiyAo=",
        version = "v28.3.3+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        sum = "h1:Dypm25kh4rmk49v1eiVbsAtpAsYURjYkaKubwuBdxEI=",
        version = "v28.3.3+incompatible",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        importpath = "github.com/docker/go-connections",
        sum = "h1:LlMG9azAe1TqfR7sO+NJttz1gy6KO7VJBh+pMmjSD94=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_docker_go_units",
        importpath = "github.com/docker/go-units",
        sum = "h1:69rxXcBk27SvSaaxTtLh/8llcHD8vYHT7WSdRZ/jvr4=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_docopt_docopt_go",
        importpath = "github.com/docopt/docopt-go",
        sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
        version = "v0.0.0-20180111231733-ee0de3bc6815",
    )
    go_repository(
        name = "com_github_dprotaso_go_yit",
        importpath = "github.com/dprotaso/go-yit",
        sum = "h1:PRxIJD8XjimM5aTknUK9w6DHLDox2r2M3DI4i2pnd3w=",
        version = "v0.0.0-20220510233725-9ba8df137936",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_dvsekhvalnov_jose2go",
        importpath = "github.com/dvsekhvalnov/jose2go",
        sum = "h1:Y9gnSnP4qEI0+/uQkHvFXeD2PLPJeXEL+ySMEA2EjTY=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_edsrzf_mmap_go",
        importpath = "github.com/edsrzf/mmap-go",
        sum = "h1:aaQcKT9WumO6JEJcRyTqFVq4XUZiUcKR2/GI31TOcz8=",
        version = "v0.0.0-20170320065105-0bce6a688712",
    )
    go_repository(
        name = "com_github_elk_language_go_prompt",
        importpath = "github.com/elk-language/go-prompt",
        sum = "h1:p6CJNCKcPUwUB4vkIvlqQNzW7ScrBHHKfMdFyeoESbc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:zEqyPVyku6IvWCFwux4x9RxkLOMUL+1vC9xUFv5l2/M=",
        version = "v0.13.4",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane_envoy",
        importpath = "github.com/envoyproxy/go-control-plane/envoy",
        sum = "h1:jb83lalDRZSpPWW2Z7Mck/8kXZ5CQAFYVjQcdVIr83A=",
        version = "v1.32.4",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane_ratelimit",
        importpath = "github.com/envoyproxy/go-control-plane/ratelimit",
        sum = "h1:/G9QYbddjL25KvtKTv3an9lx6VBE2cnb8wp1vEGNYGI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:DEo3O99U8j4hBFwbJfrz9VtgcDfUKS7KJ7spH3d86P8=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_erikgeiser_coninput",
        importpath = "github.com/erikgeiser/coninput",
        sum = "h1:Y/CXytFA4m6baUTXGLOoWe4PQhGxaX0KpnayAqC48p4=",
        version = "v0.0.0-20211004153227-1c3628e74d0f",
    )
    go_repository(
        name = "com_github_expr_lang_expr",
        importpath = "github.com/expr-lang/expr",
        sum = "h1:1h6i8ONk9cexhDmowO/A64VPxHScu7qfSl2k8OlINec=",
        version = "v1.17.6",
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:S8gINlzdQ840/4pfAwic/ZE0djQEH3wM94VfqLTZcOM=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_fatih_structs",
        importpath = "github.com/fatih/structs",
        sum = "h1:Q7juDM0QtcnhCpeyLGQKyg4TOIghuNXrkL32pHAUMxo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_fogleman_gg",
        importpath = "github.com/fogleman/gg",
        sum = "h1:/7zJX8F6AaYQc57WQCyN9cAIz+4bCJGO9B+dyW29am8=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_form3tech_oss_jwt_go",
        importpath = "github.com/form3tech-oss/jwt-go",
        sum = "h1:/l4kBbb4/vGSsdtB5nUe8L7B9mImVMaBPw9L/0TBHU8=",
        version = "v3.2.5+incompatible",
    )
    go_repository(
        name = "com_github_frankban_quicktest",
        importpath = "github.com/frankban/quicktest",
        sum = "h1:7Xjx+VpznH+oBnejlPUj8oUpdxnVs4f8XU8WnHkI4W8=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_github_fsouza_fake_gcs_server",
        importpath = "github.com/fsouza/fake-gcs-server",
        sum = "h1:OeH75kBZcZa3ZE+zz/mFdJ2btt9FgqfjI7gIh9+5fvk=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_github_fullstorydev_grpcurl",
        importpath = "github.com/fullstorydev/grpcurl",
        sum = "h1:JMvZXK8lHDGyLmTQ0ZdGDnVVGuwjbpaumf8p42z0d+c=",
        version = "v1.8.9",
    )
    go_repository(
        name = "com_github_gabriel_vasile_mimetype",
        importpath = "github.com/gabriel-vasile/mimetype",
        sum = "h1:5k+WDwEsD9eTLL8Tz3L0VnmVh9QxGjRmjBvAG7U/oYY=",
        version = "v1.4.9",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gin_contrib_cors",
        importpath = "github.com/gin-contrib/cors",
        sum = "h1:3gQ8GMzs1Ylpf70y8bMw4fVpycXIeX1ZemuSQIsnQQY=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_github_gin_contrib_sse",
        importpath = "github.com/gin-contrib/sse",
        sum = "h1:n0w2GMuUpWDVp7qSpvze6fAu9iRxJY4Hmj6AmBOU05w=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_gin_gonic_gin",
        importpath = "github.com/gin-gonic/gin",
        sum = "h1:T0ujvqyCSqRopADpgPgiTT63DUQVSfojyME59Ei63pQ=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_glebarez_go_sqlite",
        importpath = "github.com/glebarez/go-sqlite",
        sum = "h1:uAcMJhaA6r3LHMTFgP0SifzgXg46yJkgxqyuyec+ruQ=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_github_glebarez_sqlite",
        importpath = "github.com/glebarez/sqlite",
        sum = "h1:wSG0irqzP6VurnMEpFGer5Li19RpIRi2qvQz++w0GMw=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_gliderlabs_ssh",
        importpath = "github.com/gliderlabs/ssh",
        sum = "h1:a4YXD1V7xMF9g5nTkdfnja3Sxy1PVDCj1Zg4Wb8vY6c=",
        version = "v0.3.8",
    )
    go_repository(
        name = "com_github_go_asn1_ber_asn1_ber",
        importpath = "github.com/go-asn1-ber/asn1-ber",
        sum = "h1:vXT6d/FNDiELJnLb6hGNa309LMsrCoYFvpwHDF0+Y1A=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_go_fonts_dejavu",
        importpath = "github.com/go-fonts/dejavu",
        sum = "h1:JSajPXURYqpr+Cu8U9bt8K+XcACIHWqWrvWCKyeFmVQ=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_go_fonts_latin_modern",
        importpath = "github.com/go-fonts/latin-modern",
        sum = "h1:5/Tv1Ek/QCr20C6ZOz15vw3g7GELYL98KWr8Hgo+3vk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_fonts_liberation",
        importpath = "github.com/go-fonts/liberation",
        sum = "h1:jAkAWJP4S+OsrPLZM4/eC9iW7CtHy+HBXrEwZXWo5VM=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_fonts_stix",
        importpath = "github.com/go-fonts/stix",
        sum = "h1:UlZlgrvvmT/58o573ot7NFw0vZasZ5I6bcIft/oMdgg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_go_gl_glfw",
        importpath = "github.com/go-gl/glfw",
        sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
        version = "v0.0.0-20190409004039-e6da0acd62b1",
    )
    go_repository(
        name = "com_github_go_gl_glfw_v3_3_glfw",
        importpath = "github.com/go-gl/glfw/v3.3/glfw",
        sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
        version = "v0.0.0-20200222043503-6f7a984d4dc4",
    )
    go_repository(
        name = "com_github_go_jose_go_jose_v4",
        importpath = "github.com/go-jose/go-jose/v4",
        sum = "h1:TK/7NqRQZfgAh+Td8AlsrvtPoUyiHh0LqVvokh+1vHI=",
        version = "v4.1.2",
    )
    go_repository(
        name = "com_github_go_json_experiment_json",
        importpath = "github.com/go-json-experiment/json",
        sum = "h1:iizUGZ9pEquQS5jTGkh4AqeeHCMbfbjeb0zMt0aEFzs=",
        version = "v0.0.0-20250725192818-e39067aee2d2",
    )
    go_repository(
        name = "com_github_go_latex_latex",
        importpath = "github.com/go-latex/latex",
        sum = "h1:6zl3BbBhdnMkpSj2YY30qV3gDcVBGtFgVsV3+/i+mKQ=",
        version = "v0.0.0-20210823091927-c0d11ff05a81",
    )
    go_repository(
        name = "com_github_go_ldap_ldap_v3",
        importpath = "github.com/go-ldap/ldap/v3",
        sum = "h1:qPjipEpt+qDa6SI/h1fzuGWoRUY+qqQ9sOZq67/PYUs=",
        version = "v3.4.4",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        sum = "h1:CjnDlHq8ikf6E492q6eKboGOC0T8CDaOvkHCIg8idEI=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_ole_go_ole",
        importpath = "github.com/go-ole/go-ole",
        sum = "h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=",
        version = "v1.2.6",
    )
    go_repository(
        name = "com_github_go_pdf_fpdf",
        importpath = "github.com/go-pdf/fpdf",
        sum = "h1:MlgtGIfsdMEEQJr2le6b/HNr1ZlQwxyWr77r2aj2U/8=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_go_playground_assert_v2",
        importpath = "github.com/go-playground/assert/v2",
        sum = "h1:JvknZsQTYeFEAhQwI4qEt9cyV5ONwRHC+lYKSsYSR8s=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_go_playground_locales",
        importpath = "github.com/go-playground/locales",
        sum = "h1:EWaQ/wswjilfKLTECiXz7Rh+3BjFhfDFKv/oXslEjJA=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_github_go_playground_universal_translator",
        importpath = "github.com/go-playground/universal-translator",
        sum = "h1:Bcnm0ZwsGyWbCzImXv+pAJnYK9S473LQFuzCbDbfSFY=",
        version = "v0.18.1",
    )
    go_repository(
        name = "com_github_go_playground_validator_v10",
        importpath = "github.com/go-playground/validator/v10",
        sum = "h1:w8+XrWVMhGkxOaaowyKH35gFydVHOvC0/uWoy2Fzwn4=",
        version = "v10.27.0",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:U/N249h2WzJ3Ukj8SowVFjdtZKfu9vlLZxjPXV1aweo=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_go_viper_mapstructure_v2",
        importpath = "github.com/go-viper/mapstructure/v2",
        sum = "h1:EBsztssimR/CONLSZZ04E8qAkxNYq4Qp9LvH92wZUgs=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_gobuffalo_here",
        importpath = "github.com/gobuffalo/here",
        sum = "h1:hYrd0a6gDmWxBM4TnrGw8mQg24iSVoIkHEk7FodQcBI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_gobwas_httphead",
        importpath = "github.com/gobwas/httphead",
        sum = "h1:exrUm0f4YX0L7EBwZHuCF4GDp8aJfVeBrlLQrs6NqWU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobwas_pool",
        importpath = "github.com/gobwas/pool",
        sum = "h1:xfeeEhW7pwmX8nuLVlqbzVc7udMDrwetjEv+TZIz1og=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_gobwas_ws",
        importpath = "github.com/gobwas/ws",
        sum = "h1:CTaoG1tojrh4ucGPcoJFiAQUAsEWekEWvLy7GsVNqGs=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_goccmack_gocc",
        importpath = "github.com/goccmack/gocc",
        sum = "h1:FSii2UQeSLngl3jFoR4tUKZLprO7qUlh/TKKticc0BM=",
        version = "v0.0.0-20230228185258-2292f9e40198",
    )
    go_repository(
        name = "com_github_goccy_go_json",
        importpath = "github.com/goccy/go-json",
        sum = "h1:Fq85nIqj+gXn/S5ahsiTlK3TmC85qgirsdTP/+DeaC4=",
        version = "v0.10.5",
    )
    go_repository(
        name = "com_github_goccy_go_yaml",
        importpath = "github.com/goccy/go-yaml",
        sum = "h1:8W7wMFS12Pcas7KU+VVkaiCng+kG8QiFeFwzFb+rwuw=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_gocql_gocql",
        importpath = "github.com/gocql/gocql",
        sum = "h1:N/MD/sr6o61X+iZBAT2qEUF023s4KbA8RWfKzl0L6MQ=",
        version = "v0.0.0-20210515062232-b7ef815b4556",
    )
    go_repository(
        name = "com_github_godbus_dbus",
        importpath = "github.com/godbus/dbus",
        sum = "h1:ZpnhV/YsD2/4cESfV5+Hoeu/iUR3ruzNvZ+yQfO03a0=",
        version = "v0.0.0-20190726142602-4481cbc300e2",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_golang_freetype",
        importpath = "github.com/golang/freetype",
        sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
        version = "v0.0.0-20170609003504-e2365dfdc4a0",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:DrW6hGnjIhtvhOIiAKT6Psh/Kd/ldepEa81DKeiRJ5I=",
        version = "v1.2.5",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:f+oWsMOmNPc8JmEHVZIycC7hBoQxHH9pNKQORJNozsQ=",
        version = "v0.0.0-20241129210726-2c02b8208cf8",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:YtQM7lnr8iZ+j5q71MGKkNw9Mn7AjHM68uc9g5fXeUI=",
        version = "v4.5.2",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v5",
        importpath = "github.com/golang-jwt/jwt/v5",
        sum = "h1:pv4AsKCKKZuqlgs5sUmn4x8UlGa0kEVt/puTpKx9vvo=",
        version = "v5.3.0",
    )
    go_repository(
        name = "com_github_golang_migrate_migrate_v4",
        importpath = "github.com/golang-migrate/migrate/v4",
        sum = "h1:EYGkoOsvgHHfm5U/naS1RP/6PL/Xv3S4B/swMiAmDLs=",
        version = "v4.18.3",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:YojYx61/OLFsiv6Rw1Z96LpldJIy31o+UHmwAUMJ6/U=",
        version = "v1.7.0-rc.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_golang_sql_civil",
        importpath = "github.com/golang-sql/civil",
        sum = "h1:au07oEsX2xN0ktxqI+Sida1w446QrXBRJ0nee3SNZlA=",
        version = "v0.0.0-20220223132316-b832511892a9",
    )
    go_repository(
        name = "com_github_golang_sql_sqlexp",
        importpath = "github.com/golang-sql/sqlexp",
        sum = "h1:ZCD6MBpcuOVfGVqsEmY5/4FtYiKz6tSyUv9LPEDei6A=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_btree",
        importpath = "github.com/google/btree",
        sum = "h1:CVpQJjYgC4VbzxeGVHfvZrv1ctoYCAI8vbl07Fcxlyg=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_google_flatbuffers",
        importpath = "github.com/google/flatbuffers",
        sum = "h1:M9dgRyhJemaM4Sw8+66GHBu8ioaQmyPLg1b8VwK5WJg=",
        version = "v23.5.26+incompatible",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:wk8382ETsv4JYUZwIsn6YpYiWiBsYLSJiTsyBybVuN8=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_google_go_github_v39",
        importpath = "github.com/google/go-github/v39",
        sum = "h1:rNNM311XtPOz5rDdsJXAp2o8F67X9FnROXTvto3aSnQ=",
        version = "v39.2.0",
    )
    go_repository(
        name = "com_github_google_go_github_v58",
        importpath = "github.com/google/go-github/v58",
        sum = "h1:Una7GGERlF/37XfkPwpzYJe0Vp4dt2k1kCjlxwjIvzw=",
        version = "v58.0.0",
    )
    go_repository(
        name = "com_github_google_go_github_v64",
        importpath = "github.com/google/go-github/v64",
        sum = "h1:4G61sozmY3eiPAjjoOHponXDBONm+utovTKbyUb2Qdg=",
        version = "v64.0.0",
    )
    go_repository(
        name = "com_github_google_go_github_v66",
        importpath = "github.com/google/go-github/v66",
        sum = "h1:ADJsaXj9UotwdgK8/iFZtv7MLc8E8WBl62WLd/D/9+M=",
        version = "v66.0.0",
    )
    go_repository(
        name = "com_github_google_go_github_v67",
        importpath = "github.com/google/go-github/v67",
        sum = "h1:g11NDAmfaBaCO8qYdI9fsmbaRipHNWRIU/2YGvlh4rg=",
        version = "v67.0.0",
    )
    go_repository(
        name = "com_github_google_go_github_v71",
        importpath = "github.com/google/go-github/v71",
        sum = "h1:Zi16OymGKZZMm8ZliffVVJ/Q9YZreDKONCr+WUd0Z30=",
        version = "v71.0.0",
    )
    go_repository(
        name = "com_github_google_go_github_v72",
        importpath = "github.com/google/go-github/v72",
        sum = "h1:FcIO37BLoVPBO9igQQ6tStsv2asG4IPcYFi655PPvBM=",
        version = "v72.0.0",
    )
    go_repository(
        name = "com_github_google_go_pkcs11",
        importpath = "github.com/google/go-pkcs11",
        sum = "h1:PVRnTgtArZ3QQqTGtbtjtnIkzl2iY2kt24yqbrf7td8=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_google_go_querystring",
        importpath = "github.com/google/go-querystring",
        sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_google_martian",
        importpath = "github.com/google/martian",
        sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        importpath = "github.com/google/martian/v3",
        sum = "h1:DIhPTQrbPkgs2yJYdXU/eNACCG5DVQjySNRNlflZ9Fc=",
        version = "v3.3.3",
    )
    go_repository(
        name = "com_github_google_pprof",
        importpath = "github.com/google/pprof",
        sum = "h1:ijClszYn+mADRFY17kjQEVQ1XRhq2/JR1M3sGqeJoxs=",
        version = "v0.0.0-20250317173921-a4b03ec1a45e",
    )
    go_repository(
        name = "com_github_google_renameio",
        importpath = "github.com/google/renameio",
        sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_s2a_go",
        importpath = "github.com/google/s2a-go",
        sum = "h1:LGD7gtMgezd8a/Xak7mEWL0PjoTQFvpRudN895yqKW0=",
        version = "v0.1.9",
    )
    go_repository(
        name = "com_github_google_shlex",
        importpath = "github.com/google/shlex",
        sum = "h1:El6M4kTTCOh6aBiKaUGG7oYTSPP8MxqL4YI3kZKwcP4=",
        version = "v0.0.0-20191202100458-e7afc7fbc510",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_googleapis_cloud_bigtable_clients_test",
        importpath = "github.com/googleapis/cloud-bigtable-clients-test",
        sum = "h1:afMKTvA/jc6jSTMkeHBZGFDTt8Cc+kb1ATFzqMK85hw=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:GW/XbdyBFQ8Qe+YAmFU9uHLo7OnF5tL52HFAgMmyrf4=",
        version = "v0.3.6",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:SyjDc1mGgZU5LncH8gimWo9lW1DtIfPibOG81vgd/bo=",
        version = "v2.15.0",
    )
    go_repository(
        name = "com_github_googleapis_go_sql_spanner",
        importpath = "github.com/googleapis/go-sql-spanner",
        sum = "h1:VgtvMm5n4KHjH+lHmi6YYjPskCYwArL+7hKGHZj3Cgc=",
        version = "v1.16.3",
    )
    go_repository(
        name = "com_github_googleapis_go_type_adapters",
        importpath = "github.com/googleapis/go-type-adapters",
        sum = "h1:9XdMn+d/G57qq1s8dNc5IesGCXHf6V2HZ2JwRxfA2tA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_googleapis_google_cloud_go_testing",
        importpath = "github.com/googleapis/google-cloud-go-testing",
        sum = "h1:tlyzajkF3030q6M8SvmJSemC9DTHL/xaMa18b65+JM4=",
        version = "v0.0.0-20200911160855-bcd43fbb19e8",
    )
    go_repository(
        name = "com_github_googlecloudplatform_grpc_gcp_go_grpcgcp",
        importpath = "github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp",
        sum = "h1:2afWGsMzkIcN8Qm4mgPJKZWyroE5QBszMiDMYEBrnfw=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_detectors_gcp",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp",
        sum = "h1:UQUsRi8WTzhZntp5313l+CHIAT95ojUI2lpP/ExlZa4=",
        version = "v1.29.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_exporter_metric",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric",
        sum = "h1:owcC2UnmsZycprQ5RfRgjydWhuoxg71LUfyiQdijZuM=",
        version = "v0.53.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_exporter_trace",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace",
        sum = "h1:YVtMlmfRUTaWs3+1acwMBp7rBUo6zrxl6Kn13/R9YW4=",
        version = "v1.29.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_internal_cloudmock",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/cloudmock",
        sum = "h1:4LP6hvB4I5ouTbGgWtixJhgED6xdf67twf9PoY96Tbg=",
        version = "v0.53.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_internal_resourcemapping",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping",
        sum = "h1:Ron4zCA/yk6U7WOBXhTJcDpsUBG9npumK6xw2auFltQ=",
        version = "v0.53.0",
    )
    go_repository(
        name = "com_github_gorilla_css",
        importpath = "github.com/gorilla/css",
        sum = "h1:ntNaBIghp6JmvWnxbZKANoLyuXTPZ4cAMlo6RyhlbO8=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_gorilla_handlers",
        importpath = "github.com/gorilla/handlers",
        sum = "h1:0QniY0USkHQ1RGCLfKxeNHK9bkDHGRYGNDFBCS+YARg=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sum = "h1:TuBL49tXwgrFYWhqrNgrUNEY92u81SPhu7sTdzQEiWY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_gorilla_securecookie",
        importpath = "github.com/gorilla/securecookie",
        sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_sessions",
        importpath = "github.com/gorilla/sessions",
        sum = "h1:DHd3rPN5lE3Ts3D8rKkQ8x/0kqfeNmBAaiSi+o7FsgI=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway_v2",
        importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
        sum = "h1:8Tjv8EJ+pM1xP8mK6egEbD1OgnVTyacbefKhmbLhIhU=",
        version = "v2.27.2",
    )
    go_repository(
        name = "com_github_gsterjov_go_libsecret",
        importpath = "github.com/gsterjov/go-libsecret",
        sum = "h1:6rhixN/i8ZofjG1Y75iExal34USq5p+wiN1tpie8IrU=",
        version = "v0.0.0-20161001094733-a6f4afe4910c",
    )
    go_repository(
        name = "com_github_h2non_parth",
        importpath = "github.com/h2non/parth",
        sum = "h1:2VTzZjLZBgl62/EtslCrtky5vbi9dd7HrQPQIx6wqiw=",
        version = "v0.0.0-20190131123155-b4df798d6542",
    )
    go_repository(
        name = "com_github_hailocab_go_hostpool",
        importpath = "github.com/hailocab/go-hostpool",
        sum = "h1:5upAirOpQc1Q53c0bnx2ufif5kANL7bfZWcc6VJWJd8=",
        version = "v0.0.0-20160125115350-e80d13ce29ed",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_envparse",
        importpath = "github.com/hashicorp/go-envparse",
        sum = "h1:bE++6bhIsNCPLvgDZkYqo3nA+/PFI51pkrHdmPSDFPY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_uuid",
        importpath = "github.com/hashicorp/go-uuid",
        sum = "h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:0hERBMJE1eitiLkihrMvRVBYAkpHzc/J3QdDN+dAcgU=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru_v2",
        importpath = "github.com/hashicorp/golang-lru/v2",
        sum = "h1:a+bsQ5rvGLjzHuww6tVxozPZFVghXaHOwFs4luLUK2k=",
        version = "v2.0.7",
    )
    go_repository(
        name = "com_github_henvic_httpretty",
        importpath = "github.com/henvic/httpretty",
        sum = "h1:JdzGzKZBajBfnvlMALXXMVQWxWMF/ofTy8C3/OSUTxs=",
        version = "v0.0.6",
    )
    go_repository(
        name = "com_github_huandu_xstrings",
        importpath = "github.com/huandu/xstrings",
        sum = "h1:2ag3IFq9ZDANvthTwTiqSSZLjDc+BedvHPAp5tJy2TI=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_iancoleman_strcase",
        importpath = "github.com/iancoleman/strcase",
        sum = "h1:nTXanmYxhfFAMjZL34Ov6gkzEsSJZ5DbhxWjvSASxEI=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_ianlancetaylor_demangle",
        importpath = "github.com/ianlancetaylor/demangle",
        sum = "h1:mV02weKRL81bEnm8A0HT1/CAelMQDBuQIfLw8n+d6xI=",
        version = "v0.0.0-20200824232613-28f6c0f3b639",
    )
    go_repository(
        name = "com_github_inconshreveable_mousetrap",
        importpath = "github.com/inconshreveable/mousetrap",
        sum = "h1:wN+x4NVGpMsO7ErUn/mUI3vEoE6Jt13X2s0bqwp9tc8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_itchyny_gojq",
        importpath = "github.com/itchyny/gojq",
        sum = "h1:8av8eGduDb5+rvEdaOO+zQUjA04MS0m3Ps8HiD+fceg=",
        version = "v0.12.17",
    )
    go_repository(
        name = "com_github_itchyny_timefmt_go",
        importpath = "github.com/itchyny/timefmt-go",
        sum = "h1:ia3s54iciXDdzWzwaVKXZPbiXzxxnv1SPGFfM/myJ5Q=",
        version = "v0.1.6",
    )
    go_repository(
        name = "com_github_jackc_chunkreader_v2",
        importpath = "github.com/jackc/chunkreader/v2",
        sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_jackc_pgconn",
        importpath = "github.com/jackc/pgconn",
        sum = "h1:bVoTr12EGANZz66nZPkMInAV/KHD2TxH9npjXXgiB3w=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_github_jackc_pgerrcode",
        importpath = "github.com/jackc/pgerrcode",
        sum = "h1:s+4MhCQ6YrzisK6hFJUX53drDT4UsSW3DEhKn0ifuHw=",
        version = "v0.0.0-20220416144525-469b46aa5efa",
    )
    go_repository(
        name = "com_github_jackc_pgio",
        importpath = "github.com/jackc/pgio",
        sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgpassfile",
        importpath = "github.com/jackc/pgpassfile",
        sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3_v2",
        importpath = "github.com/jackc/pgproto3/v2",
        sum = "h1:1HLSx5H+tXR9pW3in3zaztoEwQYRC9SQaYUHjTSUOag=",
        version = "v2.3.3",
    )
    go_repository(
        name = "com_github_jackc_pgservicefile",
        importpath = "github.com/jackc/pgservicefile",
        sum = "h1:iCEnooe7UlwOQYpKFhBabPMi4aNAfoODPEFNiAnClxo=",
        version = "v0.0.0-20240606120523-5a60cdf6a761",
    )
    go_repository(
        name = "com_github_jackc_pgtype",
        importpath = "github.com/jackc/pgtype",
        sum = "h1:y+xUdabmyMkJLyApYuPj38mW+aAIqCe5uuBB51rH3Vw=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_jackc_pgx_v4",
        importpath = "github.com/jackc/pgx/v4",
        sum = "h1:xVpYkNR5pk5bMCZGfClbO962UIqVABcAGt7ha1s/FeU=",
        version = "v4.18.2",
    )
    go_repository(
        name = "com_github_jackc_pgx_v5",
        importpath = "github.com/jackc/pgx/v5",
        sum = "h1:JHGfMnQY+IEtGM63d+NGMjoRpysB2JBwDr5fsngwmJs=",
        version = "v5.7.5",
    )
    go_repository(
        name = "com_github_jackc_puddle_v2",
        importpath = "github.com/jackc/puddle/v2",
        sum = "h1:PR8nw+E/1w0GLuRFSmiioY6UooMp6KJv0/61nB7icHo=",
        version = "v2.2.2",
    )
    go_repository(
        name = "com_github_jaswdr_faker",
        importpath = "github.com/jaswdr/faker",
        sum = "h1:xBoz8/O6r0QAR8eEvKJZMdofxiRH+F0M/7MU9eNKhsM=",
        version = "v1.19.1",
    )
    go_repository(
        name = "com_github_jcmturner_aescts_v2",
        importpath = "github.com/jcmturner/aescts/v2",
        sum = "h1:9YKLH6ey7H4eDBXW8khjYslgyqG2xZikXP0EQFKrle8=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_dnsutils_v2",
        importpath = "github.com/jcmturner/dnsutils/v2",
        sum = "h1:lltnkeZGL0wILNvrNiVCR6Ro5PGU/SeBvVO/8c/iPbo=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_gofork",
        importpath = "github.com/jcmturner/gofork",
        sum = "h1:QH0l3hzAU1tfT3rZCnW5zXl+orbkNMMRGJfdJjHVETg=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_github_jcmturner_goidentity_v6",
        importpath = "github.com/jcmturner/goidentity/v6",
        sum = "h1:VKnZd2oEIMorCTsFBnJWbExfNN7yZr3EhJAxwOkZg6o=",
        version = "v6.0.1",
    )
    go_repository(
        name = "com_github_jcmturner_gokrb5_v8",
        importpath = "github.com/jcmturner/gokrb5/v8",
        sum = "h1:x1Sv4HaTpepFkXbt2IkL29DXRf8sOfZXo8eRKh687T8=",
        version = "v8.4.4",
    )
    go_repository(
        name = "com_github_jcmturner_rpc_v2",
        importpath = "github.com/jcmturner/rpc/v2",
        sum = "h1:7FXXj8Ti1IaVFpSAziCZWNzbNuZmnvw/i6CqLNdWfZY=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_jhump_gopoet",
        importpath = "github.com/jhump/gopoet",
        sum = "h1:gYjOPnzHd2nzB37xYQZxj4EIQNpBrBskRqQQ3q4ZgSg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_jhump_goprotoc",
        importpath = "github.com/jhump/goprotoc",
        sum = "h1:Y1UgUX+txUznfqcGdDef8ZOVlyQvnV0pKWZH08RmZuo=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_jhump_protoreflect",
        importpath = "github.com/jhump/protoreflect",
        sum = "h1:OUsOWe/nhWohrzIjKP7Wk3Bt1lhDHn0w39uiT/zTWPM=",
        version = "v1.17.1-0.20240913204751-8f5fd1dcb3c5",
    )
    go_repository(
        name = "com_github_jhump_protoreflect_v2",
        importpath = "github.com/jhump/protoreflect/v2",
        sum = "h1:qZU+rEZUOYTz1Bnhi3xbwn+VxdXkLVeEpAeZzVXLY88=",
        version = "v2.0.0-beta.2",
    )
    go_repository(
        name = "com_github_jinzhu_inflection",
        importpath = "github.com/jinzhu/inflection",
        sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jinzhu_now",
        importpath = "github.com/jinzhu/now",
        sum = "h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_johncgriffin_overflow",
        importpath = "github.com/JohnCGriffin/overflow",
        sum = "h1:RGWPOewvKIROun94nF7v2cua9qP+thov/7M50KEoeSU=",
        version = "v0.0.0-20211019200055-46fa312c352c",
    )
    go_repository(
        name = "com_github_josharian_mapfs",
        importpath = "github.com/josharian/mapfs",
        sum = "h1:c+ctPFdISggaSNCfU1IueNBAsqetJSvMcpQlT+0OVdY=",
        version = "v0.0.0-20210615234106-095c008854e6",
    )
    go_repository(
        name = "com_github_josharian_txtarfs",
        importpath = "github.com/josharian/txtarfs",
        sum = "h1:ZWuoyLMwZvLJ6OHUhPq1sZHa37Pikt6DXkZPhhOBzEE=",
        version = "v0.0.0-20240408113805-5dc76b8fe6bf",
    )
    go_repository(
        name = "com_github_jpillora_backoff",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_jstemmer_go_junit_report",
        importpath = "github.com/jstemmer/go-junit-report",
        sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_juliangruber_go_intersect",
        importpath = "github.com/juliangruber/go-intersect",
        sum = "h1:sc+y5dCjMMx0pAdYk/N6KBm00tD/f3tq+Iox7dYDUrY=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_jung_kurt_gofpdf",
        importpath = "github.com/jung-kurt/gofpdf",
        sum = "h1:PJr+ZMXIecYc1Ey2zucXdR73SMBtgjPgwa31099IMv0=",
        version = "v1.0.3-0.20190309125859-24315acbbda5",
    )
    go_repository(
        name = "com_github_k0kubun_pp",
        importpath = "github.com/k0kubun/pp",
        sum = "h1:EKhKbi34VQDWJtq+zpsKSEhkHHs9w2P8Izbq8IhLVSo=",
        version = "v2.3.0+incompatible",
    )
    go_repository(
        name = "com_github_k0kubun_pp_v3",
        importpath = "github.com/k0kubun/pp/v3",
        sum = "h1:iYNlYA5HJAJvkD4ibuf9c8y6SHM0QFhaBuCqm1zHp0w=",
        version = "v3.5.0",
    )
    go_repository(
        name = "com_github_k1low_bufresolv",
        importpath = "github.com/k1LoW/bufresolv",
        sum = "h1:E1cd7z9xYtVa6Cg5VczA36Dica0/s1rBVtZfqR0brTU=",
        version = "v0.7.9",
    )
    go_repository(
        name = "com_github_k1low_concgroup",
        importpath = "github.com/k1LoW/concgroup",
        sum = "h1:h+rGuaRZnCYkMGCwFurawWIyaDNy6y0qm5RqhqlOHn8=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_k1low_curlreq",
        importpath = "github.com/k1LoW/curlreq",
        sum = "h1:ycHRTnphAq3vWZU49EdmXYYuF1AGkJBY5MQADvE7Ezo=",
        version = "v0.3.3",
    )
    go_repository(
        name = "com_github_k1low_donegroup",
        importpath = "github.com/k1LoW/donegroup",
        sum = "h1:T3XWToJqsT9Gr+WBnxKV8FupE6x7LL0UsfMiYJrD8WA=",
        version = "v1.10.2",
    )
    go_repository(
        name = "com_github_k1low_duration",
        importpath = "github.com/k1LoW/duration",
        sum = "h1:qq1gWtPh7YROFyerBufVP+ATR11mOOHDInrcC/Xe/6A=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_k1low_exec",
        importpath = "github.com/k1LoW/exec",
        sum = "h1:Wc01vrKXOAa1HfIRiDWcn3p2ebl2qVk+kOLqL7mYBL0=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_k1low_expand",
        importpath = "github.com/k1LoW/expand",
        sum = "h1:pPT/BargN78RScVGaOJCGZo6t0BIxOJXWLlkUq1T+C4=",
        version = "v0.16.2",
    )
    go_repository(
        name = "com_github_k1low_ghfs",
        importpath = "github.com/k1LoW/ghfs",
        sum = "h1:azrPgoXEMjpuDcTLc8YFV73kHWDC3MPAezO3eVWIU/E=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_k1low_go_github_client_v58",
        importpath = "github.com/k1LoW/go-github-client/v58",
        sum = "h1:h6Iaqg5aWqh+TNY7BSdumztw6id8MmL/a2bzYdeZMoU=",
        version = "v58.0.18",
    )
    go_repository(
        name = "com_github_k1low_go_github_client_v67",
        importpath = "github.com/k1LoW/go-github-client/v67",
        sum = "h1:PjzG1zRpA+kKQEq3Z0fZ/8UMJLHaLGoHptMpOPViptw=",
        version = "v67.0.13",
    )
    go_repository(
        name = "com_github_k1low_grpcstub",
        importpath = "github.com/k1LoW/grpcstub",
        sum = "h1:rlSp/xYs2G/ebv76RjS+JZeoYqVWZofYaWkkj+AFFYQ=",
        version = "v0.25.11",
    )
    go_repository(
        name = "com_github_k1low_grpcurlreq",
        importpath = "github.com/k1LoW/grpcurlreq",
        sum = "h1:vXpq94BJp57iKqHgrSk85o54Dpx4qxsrIs3axJaLi7k=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_k1low_httpstub",
        importpath = "github.com/k1LoW/httpstub",
        sum = "h1:jMC2MDY4ajJpd1HE59PiB6ze9D7GoXQr2SM1zx2Ali0=",
        version = "v0.22.0",
    )
    go_repository(
        name = "com_github_k1low_maskedio",
        importpath = "github.com/k1LoW/maskedio",
        sum = "h1:U3JNg1jVNfO+4ZEVy6cloFwgBIvvYbODmiFLIFEgjpo=",
        version = "v0.4.4",
    )
    go_repository(
        name = "com_github_k1low_protoresolv",
        importpath = "github.com/k1LoW/protoresolv",
        sum = "h1:jlKBERyIM20s/S9pk2DbOnqrSU5hpGSTA7qG5jV0Oqc=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_k1low_repin",
        importpath = "github.com/k1LoW/repin",
        sum = "h1:xcNuBBc/ISHUNBzjXNTCux4OYZND5ZMiyz4SrRtpDhg=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_k1low_runn",
        importpath = "github.com/k1LoW/runn",
        sum = "h1:iQD9FaJWhi7+K8mIRhGUKHrCvS1j68MIVUw0N5BEw+U=",
        version = "v0.135.0",
    )
    go_repository(
        name = "com_github_k1low_sshc_v4",
        importpath = "github.com/k1LoW/sshc/v4",
        sum = "h1:63mIGIMIt/EvA9MZv86xL+eQPXZb+G7mBv2pDuTBmX8=",
        version = "v4.2.1",
    )
    go_repository(
        name = "com_github_k1low_stopw",
        importpath = "github.com/k1LoW/stopw",
        sum = "h1:Y1DDVtOLYZ6gBHBJ9B5TmEYW5hG8LSr8u1/LyzxKTxg=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_github_k1low_urlfilepath",
        importpath = "github.com/k1LoW/urlfilepath",
        sum = "h1:JU2FJISuw9oGHy0SAC85O85pnYS3/Z2r0TLlIpy215E=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_k1low_waitmap",
        importpath = "github.com/k1LoW/waitmap",
        sum = "h1:FPj+AvJXfEAgVqO7s9XrGhWUGCQ20/9Ejn5L6F8Oco8=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_kardianos_osext",
        importpath = "github.com/kardianos/osext",
        sum = "h1:iQTw/8FWTuc7uiaSepXwyf3o52HaUYcV+Tu66S3F5GA=",
        version = "v0.0.0-20190222173326-2bc1f35cddc0",
    )
    go_repository(
        name = "com_github_kballard_go_shellquote",
        importpath = "github.com/kballard/go-shellquote",
        sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
        version = "v0.0.0-20180428030007-95032a82bc51",
    )
    go_repository(
        name = "com_github_kevinburke_ssh_config",
        importpath = "github.com/kevinburke/ssh_config",
        sum = "h1:x584FjTGwHzMwvHx18PXxbBVzfnxogHaAReU4gf13a4=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_klauspost_asmfmt",
        importpath = "github.com/klauspost/asmfmt",
        sum = "h1:4Ri7ox3EwapiOjCki+hw14RyKk201CN4rzyCJRFLpK4=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        sum = "h1:c/Cqfb0r+Yi+JtIEq73FWXVkRonBlf0CRNYc8Zttxdo=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_klauspost_cpuid_v2",
        importpath = "github.com/klauspost/cpuid/v2",
        sum = "h1:S4CRMLnYUhGeDFDqkGriYKdfoFlDnMtqTiI/sFzhA9Y=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_kr_fs",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_ktrysmt_go_bitbucket",
        importpath = "github.com/ktrysmt/go-bitbucket",
        sum = "h1:C8dUGp0qkwncKtAnozHCbbqhptefzEd1I0sfnuy9rYQ=",
        version = "v0.6.4",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_leaanthony_go_ansi_parser",
        importpath = "github.com/leaanthony/go-ansi-parser",
        sum = "h1:xd8bzARK3dErqkPFtoF9F3/HgN8UQk0ed1YDKpEz01A=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_github_ledongthuc_pdf",
        importpath = "github.com/ledongthuc/pdf",
        sum = "h1:6Yzfa6GP0rIo/kULo2bwGEkFvCePZ3qHDDTC3/J9Swo=",
        version = "v0.0.0-20220302134840-0c2507a12d80",
    )
    go_repository(
        name = "com_github_leodido_go_urn",
        importpath = "github.com/leodido/go-urn",
        sum = "h1:WT9HwE9SGECu3lg4d/dIA+jxlljEa1/ffXKmRjqdmIQ=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_lestrrat_go_backoff_v2",
        importpath = "github.com/lestrrat-go/backoff/v2",
        sum = "h1:oNb5E5isby2kiro9AgdHLv5N5tint1AnDVVf2E2un5A=",
        version = "v2.0.8",
    )
    go_repository(
        name = "com_github_lestrrat_go_blackmagic",
        importpath = "github.com/lestrrat-go/blackmagic",
        sum = "h1:lS5Zts+5HIC/8og6cGHb0uCcNCa3OUt1ygh3Qz2Fe80=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_lestrrat_go_httpcc",
        importpath = "github.com/lestrrat-go/httpcc",
        sum = "h1:ydWCStUeJLkpYyjLDHihupbn2tYmZ7m22BGkcvZZrIE=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_lestrrat_go_iter",
        importpath = "github.com/lestrrat-go/iter",
        sum = "h1:gMXo1q4c2pHmC3dn8LzRhJfP1ceCbgSiT9lUydIzltI=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_lestrrat_go_jwx",
        importpath = "github.com/lestrrat-go/jwx",
        sum = "h1:tAx93jN2SdPvFn08fHNAhqFJazn5mBBOB8Zli0g0otA=",
        version = "v1.2.25",
    )
    go_repository(
        name = "com_github_lestrrat_go_option",
        importpath = "github.com/lestrrat-go/option",
        sum = "h1:oAzP2fvZGQKWkvHa1/SAcFolBEca1oN+mQ7eooNBEYU=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_lib_pq",
        importpath = "github.com/lib/pq",
        sum = "h1:YXG7RB+JIjhP29X+OtkiDnYaXQwpS4JEWq7dtCCRUEw=",
        version = "v1.10.9",
    )
    go_repository(
        name = "com_github_lucasb_eyer_go_colorful",
        importpath = "github.com/lucasb-eyer/go-colorful",
        sum = "h1:1nnpGOrhyZZuNyfu1QjKiUICQ74+3FNCN69Aj6K7nkY=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_lucasjones_reggen",
        importpath = "github.com/lucasjones/reggen",
        sum = "h1:w1g9wNDIE/pHSTmAaUhv4TZQuPBS6GV3mMz5hkgziIU=",
        version = "v0.0.0-20200904144131-37ba4fa293bb",
    )
    go_repository(
        name = "com_github_lufia_plan9stats",
        importpath = "github.com/lufia/plan9stats",
        sum = "h1:V53FWzU6KAZVi1tPp5UIsMoUWJ2/PNwYIDXnu7QuBCE=",
        version = "v0.0.0-20230110061619-bbe2e5e100de",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        importpath = "github.com/lyft/protoc-gen-star",
        sum = "h1:erE0rdztuaDq3bpGifD95wfoPrSZc95nGA6tbiNYh6M=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star_v2",
        importpath = "github.com/lyft/protoc-gen-star/v2",
        sum = "h1:sIXJOMrYnQZJu7OB7ANSF4MYri2fTEGIsRLz6LwI4xE=",
        version = "v2.0.4-0.20230330145011-496ad1ac90a4",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
        version = "v0.7.7",
    )
    go_repository(
        name = "com_github_makenowjust_heredoc",
        importpath = "github.com/MakeNowJust/heredoc",
        sum = "h1:cXCdzVdstXyiTqTvfqk9SDHpKNjxuom+DOlyEeQ4pzQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_markbates_pkger",
        importpath = "github.com/markbates/pkger",
        sum = "h1:3MPelV53RnGSW07izx5xGxl4e/sdRD6zqseIk0rMASY=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_masterminds_goutils",
        importpath = "github.com/Masterminds/goutils",
        sum = "h1:5nUrii3FMTL5diU80unEVvNevw1nH4+ZV4DSLVJLSYI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_masterminds_semver_v3",
        importpath = "github.com/Masterminds/semver/v3",
        sum = "h1:B8LGeaivUe71a5qox1ICM/JLl0NqZSW5CHyL+hmvYS0=",
        version = "v3.3.0",
    )
    go_repository(
        name = "com_github_masterminds_sprig_v3",
        importpath = "github.com/Masterminds/sprig/v3",
        sum = "h1:mQh0Yrg1XPo6vjYXgtf5OtijNAKJRNcTdOOGZe3tPhs=",
        version = "v3.3.0",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:9A9LHSqF/7dyVVX6g0U9cwm9pG3kP9gSzcuIPHPsaIE=",
        version = "v0.1.14",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=",
        version = "v0.0.20",
    )
    go_repository(
        name = "com_github_mattn_go_localereader",
        importpath = "github.com/mattn/go-localereader",
        sum = "h1:ygSAOl7ZXTx4RdPYinUpg6W99U8jWvWi9Ye2JC/oIi4=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:E5ScNMtiwvlvB5paMFdw9p4kSQzbXFikJ5SQO6TULQc=",
        version = "v0.0.16",
    )
    go_repository(
        name = "com_github_mattn_go_shellwords",
        importpath = "github.com/mattn/go-shellwords",
        sum = "h1:M2zGm7EW6UQJvDeQxo4T51eKPurbeFbe8WtebGE2xrk=",
        version = "v1.0.12",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:2gZY6PC6kBnID23Tichd1K+Z0oS6nE/XwU+Vz/5o4kU=",
        version = "v1.14.22",
    )
    go_repository(
        name = "com_github_mattn_go_tty",
        importpath = "github.com/mattn/go-tty",
        sum = "h1:KJ486B6qI8+wBO7kQxYgmmEFDaFEE96JMBQ7h400N8Q=",
        version = "v0.0.7",
    )
    go_repository(
        name = "com_github_mgutz_ansi",
        importpath = "github.com/mgutz/ansi",
        sum = "h1:5PJl274Y63IEHC+7izoQE9x6ikvDFZS2mDVS3drnohI=",
        version = "v0.0.0-20200706080929-d51e80ef957d",
    )
    go_repository(
        name = "com_github_micahparks_keyfunc",
        importpath = "github.com/MicahParks/keyfunc",
        sum = "h1:lhKd5xrFHLNOWrDc4Tyb/Q1AJ4LCzQ48GVJyVIID3+o=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_microcosm_cc_bluemonday",
        importpath = "github.com/microcosm-cc/bluemonday",
        sum = "h1:MpEUotklkwCSLeH+Qdx1VJgNqLlpY2KXwXFM08ygZfk=",
        version = "v1.0.27",
    )
    go_repository(
        name = "com_github_microsoft_go_mssqldb",
        importpath = "github.com/microsoft/go-mssqldb",
        sum = "h1:nY8TmFMQOHpm2qVWo6y4I2mAmVdZqlGiMGAYt64Ibbs=",
        version = "v1.9.2",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        importpath = "github.com/Microsoft/go-winio",
        sum = "h1:F2VQgta7ecxGYO8k3ZZz3RS8fVIXVxONVUPlNERoyfY=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_migueleliasweb_go_github_mock",
        importpath = "github.com/migueleliasweb/go-github-mock",
        sum = "h1:2sVP9JEMB2ubQw1IKto3/fzF51oFC6eVWOOFDgQoq88=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_minio_asm2plan9s",
        importpath = "github.com/minio/asm2plan9s",
        sum = "h1:AMFGa4R4MiIpspGNG7Z948v4n35fFGB3RR3G/ry4FWs=",
        version = "v0.0.0-20200509001527-cdd76441f9d8",
    )
    go_repository(
        name = "com_github_minio_c2goasm",
        importpath = "github.com/minio/c2goasm",
        sum = "h1:+n/aFZefKZp7spd8DFdX7uMikMLXX4oubIzJF4kv/wI=",
        version = "v0.0.0-20190812172519-36a3d3bbc4f3",
    )
    go_repository(
        name = "com_github_minio_madmin_go_v3",
        importpath = "github.com/minio/madmin-go/v3",
        sum = "h1:+WuNw0q8gYTNHUmV5X1nCox28uYmJkeMT75vh9VKPkA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "com_github_minio_md5_simd",
        importpath = "github.com/minio/md5-simd",
        sum = "h1:Gdi1DZK69+ZVMoNHRXJyNcxrMA4dSxoYHZSQbirFg34=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_minio_minio_go_v7",
        importpath = "github.com/minio/minio-go/v7",
        sum = "h1:dE5DfOtnXMXCjr/HWI6zN9vCrY6Sv666qhhiwUMvGV4=",
        version = "v7.0.49",
    )
    go_repository(
        name = "com_github_minio_mux",
        importpath = "github.com/minio/mux",
        sum = "h1:r9oVDFM09y+u8CF4HPLanguAG41niXgYwZAFkVHce9M=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_minio_pkg",
        importpath = "github.com/minio/pkg",
        sum = "h1:UOUJjewE5zoaDPlCMJtNx/swc1jT1ZR+IajT7hrLd44=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_github_mitchellh_copystructure",
        importpath = "github.com/mitchellh/copystructure",
        sum = "h1:vpKXTN4ewci03Vljg/q9QvCGUDttBOGBIa15WveJJGw=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_mitchellh_reflectwalk",
        importpath = "github.com/mitchellh/reflectwalk",
        sum = "h1:G2LzWKi524PWgd3mLHV8Y5k7s6XUvT0Gef6zxSIeXaQ=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_moby_docker_image_spec",
        importpath = "github.com/moby/docker-image-spec",
        sum = "h1:jMKff3w6PgbfSa69GfNg+zN/XLhfXJGnEx3Nl2EsFP0=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_moby_sys_user",
        importpath = "github.com/moby/sys/user",
        sum = "h1:jhcMKit7SA80hivmFJcbB1vqmw//wU61Zdui2eQXuMs=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_moby_term",
        importpath = "github.com/moby/term",
        sum = "h1:6qk3FJAFDs6i/q3W/pQ97SX192qKfZgGjCQqfCJkgzQ=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_modocache_gover",
        importpath = "github.com/modocache/gover",
        sum = "h1:8Q0qkMVC/MmWkpIdlvZgcv2o2jrlF6zqVOh7W5YHdMA=",
        version = "v0.0.0-20171022184752-b58185e213c5",
    )
    go_repository(
        name = "com_github_montanaflynn_stats",
        importpath = "github.com/montanaflynn/stats",
        sum = "h1:r3y12KyNxj/Sb/iOE46ws+3mS1+MZca1wlHQFPsY/JU=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        importpath = "github.com/morikuni/aec",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mtibben_percent",
        importpath = "github.com/mtibben/percent",
        sum = "h1:5gssi8Nqo8QU/r2pynCm+hBQHpkB/uNK7BJCFogWdzs=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_muesli_ansi",
        importpath = "github.com/muesli/ansi",
        sum = "h1:ZK8zHtRHOkbHy6Mmr5D264iyp3TiX5OmNcI5cIARiQI=",
        version = "v0.0.0-20230316100256-276c6243b2f6",
    )
    go_repository(
        name = "com_github_muesli_cancelreader",
        importpath = "github.com/muesli/cancelreader",
        sum = "h1:3I4Kt4BQjOR54NavqnDogx/MIoWBFa0StPA8ELUXHmA=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_muesli_reflow",
        importpath = "github.com/muesli/reflow",
        sum = "h1:IFsN6K9NfGtjeggFP+68I4chLZV2yIKsXJFNZ+eWh6s=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_muesli_termenv",
        importpath = "github.com/muesli/termenv",
        sum = "h1:S5AlUN9dENB57rsbnkPyfdGuWIlkmzJjbFf0Tf5FWUc=",
        version = "v0.16.0",
    )
    go_repository(
        name = "com_github_munnerz_goautoneg",
        importpath = "github.com/munnerz/goautoneg",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
    )
    go_repository(
        name = "com_github_mutecomm_go_sqlcipher_v4",
        importpath = "github.com/mutecomm/go-sqlcipher/v4",
        sum = "h1:sV1tWCWGAVlPhNGT95Q+z/txFxuhAYWwHD1afF5bMZg=",
        version = "v4.4.0",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_nakagami_firebirdsql",
        importpath = "github.com/nakagami/firebirdsql",
        sum = "h1:P48LjvUQpTReR3TQRbxSeSBsMXzfK0uol7eRcr7VBYQ=",
        version = "v0.0.0-20190310045651-3c02a58cfed8",
    )
    go_repository(
        name = "com_github_ncruces_go_strftime",
        importpath = "github.com/ncruces/go-strftime",
        sum = "h1:bY0MQC28UADQmHmaF5dgpLmImcShSi2kHU9XLdhx/f4=",
        version = "v0.1.9",
    )
    go_repository(
        name = "com_github_neo4j_neo4j_go_driver",
        importpath = "github.com/neo4j/neo4j-go-driver",
        sum = "h1:fhFP5RliM2HW/8XdcO5QngSfFli9GcRIpMXvypTQt6E=",
        version = "v1.8.1-0.20200803113522-b626aa943eba",
    )
    go_repository(
        name = "com_github_niemeyer_pretty",
        importpath = "github.com/niemeyer/pretty",
        sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
        version = "v0.0.0-20200227124842-a10e7caefd8e",
    )
    go_repository(
        name = "com_github_nvveen_gotty",
        importpath = "github.com/Nvveen/Gotty",
        sum = "h1:TngWCqHvy9oXAN6lEVMRuU21PR1EtLVZJmdB18Gu3Rw=",
        version = "v0.0.0-20120604004816-cd527374f1e5",
    )
    go_repository(
        name = "com_github_ohler55_ojg",
        importpath = "github.com/ohler55/ojg",
        sum = "h1:njM65m+ej8sLHiFZIhJK9UkwOmDPsUikjGbTgcwu8CU=",
        version = "v1.26.8",
    )
    go_repository(
        name = "com_github_oklog_ulid_v2",
        importpath = "github.com/oklog/ulid/v2",
        sum = "h1:suPZ4ARWLOJLegGFiZZ1dFAkqzhMjL3J1TzI+5wHz8s=",
        version = "v2.1.1",
    )
    go_repository(
        name = "com_github_olekukonko_errors",
        importpath = "github.com/olekukonko/errors",
        sum = "h1:RNuGIh15QdDenh+hNvKrJkmxxjV4hcS50Db478Ou5sM=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_olekukonko_ll",
        importpath = "github.com/olekukonko/ll",
        sum = "h1:Y+1YqDfVkqMWuEQMclsF9HUR5+a82+dxJuL1HHSRpxI=",
        version = "v0.0.9",
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:XGwRsYLC2bY7bNd93Dk51bcPZksWZmLYuaTHR0FqfL8=",
        version = "v1.0.9",
    )
    go_repository(
        name = "com_github_olekukonko_ts",
        importpath = "github.com/olekukonko/ts",
        sum = "h1:LiZB1h0GIcudcDci2bxbqI6DXV8bF8POAnArqvRrIyw=",
        version = "v0.0.0-20171002115256-78ecb04241c0",
    )
    go_repository(
        name = "com_github_oneofone_xxhash",
        importpath = "github.com/OneOfOne/xxhash",
        sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:29JGrr5oVBm5ulCWet69zQkzWipVXIol6ygQUe/EzNc=",
        version = "v1.16.4",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        sum = "h1:WjP/FQ/sk43MRmnEcT+MlDw2TFvkrXlprrPST/IudjU=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        importpath = "github.com/opencontainers/go-digest",
        sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        importpath = "github.com/opencontainers/image-spec",
        sum = "h1:y0fUlFfIZhPF1W537XOLg0/fcx6zcHCJwooC2xJA040=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        importpath = "github.com/opencontainers/runc",
        sum = "h1:cvP7xbEvD0QQAs0nZKLzkVog2OPZhI/V2w3WmTmUSXI=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_orandin_slog_gorm",
        importpath = "github.com/orandin/slog-gorm",
        sum = "h1:FgA8hJufF9/jeNSYoEXmHPPBwET2gwlF3B85JdpsTUU=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_orisano_pixelmatch",
        importpath = "github.com/orisano/pixelmatch",
        sum = "h1:x0TT0RDC7UhAVbbWWBzr41ElhJx5tXPWkIHA2HWPRuw=",
        version = "v0.0.0-20220722002657-fb0b55479cde",
    )
    go_repository(
        name = "com_github_ory_dockertest_v3",
        importpath = "github.com/ory/dockertest/v3",
        sum = "h1:3oV9d0sDzlSQfHtIaB5k6ghUCVMVLpAY8hwrqoCyRCw=",
        version = "v3.12.0",
    )
    go_repository(
        name = "com_github_pb33f_libopenapi",
        importpath = "github.com/pb33f/libopenapi",
        sum = "h1:B0rf9Reo63tAx54gpoP9778Y84gk3JQoFtj7yg8vKfo=",
        version = "v0.25.3",
    )
    go_repository(
        name = "com_github_pb33f_libopenapi_validator",
        importpath = "github.com/pb33f/libopenapi-validator",
        sum = "h1:sS6RvphkhlgMdad4WutRVd/yzNu/7QE4RdUTjxp0dY4=",
        version = "v0.4.7",
    )
    go_repository(
        name = "com_github_pb33f_ordered_map_v2",
        importpath = "github.com/pb33f/ordered-map/v2",
        sum = "h1:+6D6e0nkcEjVPh6kF48ynz2Cb+D/ECH/Q3AOunHtj7E=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_pbnjay_memory",
        importpath = "github.com/pbnjay/memory",
        sum = "h1:onHthvaw9LFnH4t2DcNVpwGmV9E1BkGknEliJkfwQj0=",
        version = "v0.0.0-20210728143218-7b4eea64cf58",
    )
    go_repository(
        name = "com_github_pborman_getopt",
        importpath = "github.com/pborman/getopt",
        sum = "h1:BHT1/DKsYDGkUgQ2jmMaozVcdk+sVfz0+1ZJq4zkWgw=",
        version = "v0.0.0-20170112200414-7148bc3a4c30",
    )
    go_repository(
        name = "com_github_pelletier_go_toml_v2",
        importpath = "github.com/pelletier/go-toml/v2",
        sum = "h1:mye9XuhQ6gvn5h28+VilKrrPoQVanw5PMw/TB0t5Ec4=",
        version = "v2.2.4",
    )
    go_repository(
        name = "com_github_philhofer_fwd",
        importpath = "github.com/philhofer/fwd",
        sum = "h1:bnDivRJ1EWPjUIRXV5KfORO897HTbpFAQddBdE8t7Gw=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_phpdave11_gofpdf",
        importpath = "github.com/phpdave11/gofpdf",
        sum = "h1:KPKiIbfwbvC/wOncwhrpRdXVj2CZTCFlw4wnoyjtHfQ=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_phpdave11_gofpdi",
        importpath = "github.com/phpdave11/gofpdi",
        sum = "h1:o61duiW8M9sMlkVXWlvP92sZJtGKENvW3VExs6dZukQ=",
        version = "v1.0.13",
    )
    go_repository(
        name = "com_github_pierrec_lz4_v4",
        importpath = "github.com/pierrec/lz4/v4",
        sum = "h1:xaKrnTkyoqfh1YItXl56+6KJNVYWlEEPuAQW9xsplYQ=",
        version = "v4.1.18",
    )
    go_repository(
        name = "com_github_pkg_browser",
        importpath = "github.com/pkg/browser",
        sum = "h1:+mdjkGKdHQG3305AYmdv1U2eRNDiU2ErMBj1gwrq8eQ=",
        version = "v0.0.0-20240102092130-5ac0b6a4141c",
    )
    go_repository(
        name = "com_github_pkg_diff",
        importpath = "github.com/pkg/diff",
        sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
        version = "v0.0.0-20210226163009-20ebb0f2a09e",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        importpath = "github.com/pkg/sftp",
        sum = "h1:I2qBYMChEhIjOgazfJmV3/mZM256btk6wkCDRmW7JYs=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_github_pkg_term",
        importpath = "github.com/pkg/term",
        sum = "h1:L3y/h2jkuBVFdWiJvNfYfKmzcCnILw7mJWm2JQuMppw=",
        version = "v1.2.0-beta.2",
    )
    go_repository(
        name = "com_github_planetscale_vtprotobuf",
        importpath = "github.com/planetscale/vtprotobuf",
        sum = "h1:S1hI5JiKP7883xBzZAr1ydcxrKNSVNm7+3+JwjxZEsg=",
        version = "v0.6.1-0.20250313105119-ba97887b0a25",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:Jamvg5psRIccs7FGNTlIRMkT8wgtp5eCXdBlqhYGL6U=",
        version = "v1.0.1-0.20181226105442-5d4384ee4fb2",
    )
    go_repository(
        name = "com_github_power_devops_perfstat",
        importpath = "github.com/power-devops/perfstat",
        sum = "h1:0LFwY6Q3gMACTjAbMZBjXAqTOzOwFaj2Ld6cjeQ7Rig=",
        version = "v0.0.0-20221212215047-62379fc7944b",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:ust4zpdl9r4trLY/gSjlm07PuiBq2ynaXXlptpfy8Uc=",
        version = "v1.23.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:oBsgwpGs7iVziMvrGhE53c/GrLUsZdHnqNwqPLxwZyk=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sum = "h1:QDwzd+G1twt//Kwj/Ww6E9FQq1iVMmODnILtW1t2VzE=",
        version = "v0.65.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:FuLQ+05u4ZI+SS/w9+BWEM2TXiHKsUQ9TADiRH7DuK0=",
        version = "v0.17.0",
    )
    go_repository(
        name = "com_github_remychantenay_slog_otel",
        importpath = "github.com/remychantenay/slog-otel",
        sum = "h1:xoM41ayLff2U8zlK5PH31XwD7Lk3W9wKfl4+RcmKom4=",
        version = "v1.3.4",
    )
    go_repository(
        name = "com_github_remyoudompheng_bigfft",
        importpath = "github.com/remyoudompheng/bigfft",
        sum = "h1:W09IVJc94icq4NjY3clb7Lk8O1qJ8BdBEF8z0ibU0rE=",
        version = "v0.0.0-20230129092748-24d4a6f8daec",
    )
    go_repository(
        name = "com_github_rivo_uniseg",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:WUdvkW8uEhrYfLC4ZzdpI2ztxP1I582+49Oc5Mq64VQ=",
        version = "v0.4.7",
    )
    go_repository(
        name = "com_github_rjeczalik_notify",
        importpath = "github.com/rjeczalik/notify",
        sum = "h1:6rJAzHTGKXGj76sbRgDiDcYj/HniypXmSJo1SWakZeY=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:KvO1DLK/DRN07sQ1LQKScxyZJuNnedQ5/wKSR38lUII=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_github_rqlite_gorqlite",
        importpath = "github.com/rqlite/gorqlite",
        sum = "h1:V7x0hCAgL8lNGezuex1RW1sh7VXXCqfw8nXZti66iFg=",
        version = "v0.0.0-20230708021416-2acd02b70b79",
    )
    go_repository(
        name = "com_github_rs_xid",
        importpath = "github.com/rs/xid",
        sum = "h1:fV591PaemRlL6JfRxGDEPl69wICngIQ3shQtzfy2gxU=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        sum = "h1:KqfZb0pUVN2lYqZUYRddxF4OR8ZMURnJIG5Y3VRLtww=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_ruudk_golang_pdf417",
        importpath = "github.com/ruudk/golang-pdf417",
        sum = "h1:K1Xf3bKttbF+koVGaX5xngRIZ5bVjbmPnaxE/dR08uY=",
        version = "v0.0.0-20201230142125-a7e3863a1245",
    )
    go_repository(
        name = "com_github_ryo_yamaoka_otchkiss",
        importpath = "github.com/ryo-yamaoka/otchkiss",
        sum = "h1:51Joo9YNp5i0ElSTvtpj0u7AAGvtvwy6njLYymqJTEI=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_samber_lo",
        importpath = "github.com/samber/lo",
        sum = "h1:kysRYLbHy/MB7kQZf5DSN50JHmMsNEdeY24VzJFu7wI=",
        version = "v1.51.0",
    )
    go_repository(
        name = "com_github_samber_slog_gin",
        importpath = "github.com/samber/slog-gin",
        sum = "h1:EHvZTCdLAVIaqu5czMs7u2ba/65e19C5hbT1bLTTj6c=",
        version = "v1.16.1",
    )
    go_repository(
        name = "com_github_santhosh_tekuri_jsonschema_v5",
        importpath = "github.com/santhosh-tekuri/jsonschema/v5",
        sum = "h1:lZUw3E0/J3roVtGQ+SCrUrg3ON6NgVqpn3+iol9aGu4=",
        version = "v5.3.1",
    )
    go_repository(
        name = "com_github_santhosh_tekuri_jsonschema_v6",
        importpath = "github.com/santhosh-tekuri/jsonschema/v6",
        sum = "h1:PKK9DyHxif4LZo+uQSgXNqs0jj5+xZwwfKHgph2lxBw=",
        version = "v6.0.1",
    )
    go_repository(
        name = "com_github_scaleft_sshkeys",
        importpath = "github.com/ScaleFT/sshkeys",
        sum = "h1:Yqd0cKA5PUvwV0dgRI67BDHGTsMHtGQBZbLXh1dthmE=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_scylladb_termtables",
        importpath = "github.com/scylladb/termtables",
        sum = "h1:8qmTC5ByIXO3GP/IzBkxcZ/99VITvnIETDhdFz/om7A=",
        version = "v0.0.0-20191203121021-c4c0b6d42ff4",
    )
    go_repository(
        name = "com_github_secure_io_sio_go",
        importpath = "github.com/secure-io/sio-go",
        sum = "h1:dNvY9awjabXTYGsTF1PiCySl9Ltofk9GA3VdWlo7rRc=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_shirou_gopsutil_v3",
        importpath = "github.com/shirou/gopsutil/v3",
        sum = "h1:PAWSuiAszn7IhPMBtXsbSCafej7PqUOvY6YywlQUExU=",
        version = "v3.23.2",
    )
    go_repository(
        name = "com_github_shopspring_decimal",
        importpath = "github.com/shopspring/decimal",
        sum = "h1:bxl37RwXBklmTi0C79JfXCEBD1cqqHt0bbgBAGFp81k=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_snowflakedb_gosnowflake",
        importpath = "github.com/snowflakedb/gosnowflake",
        sum = "h1:KSHXrQ5o7uso25hNIzi/RObXtnSGkFgie91X82KcvMY=",
        version = "v1.6.19",
    )
    go_repository(
        name = "com_github_songmu_axslogparser",
        importpath = "github.com/Songmu/axslogparser",
        sum = "h1:cCBU44fFED0XgXDi2OqNycLeJ8JAK0//NoGYwgOj2FQ=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_songmu_go_ltsv",
        importpath = "github.com/Songmu/go-ltsv",
        sum = "h1:veR1K9TBM0PiGpxKobcJg78uiZw/FPlStpgnCHe+4tQ=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_songmu_prompter",
        importpath = "github.com/Songmu/prompter",
        sum = "h1:IAsttKsOZWSDw7bV1mtGn9TAmLFAjXbp9I/eYmUUogo=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_spaolacci_murmur3",
        importpath = "github.com/spaolacci/murmur3",
        sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
        version = "v0.0.0-20180118202830-f09979ecbc72",
    )
    go_repository(
        name = "com_github_speakeasy_api_jsonpath",
        importpath = "github.com/speakeasy-api/jsonpath",
        sum = "h1:Mys71yd6u8kuowNCR0gCVPlVAHCmKtoGXYoAtcEbqXQ=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:EaGW2JJh15aKOejeuJ+wpFSHnbd7GE6Wvp3TsNhb6LY=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        sum = "h1:SsGfm7M8QOFtEzumm7UZrZdLLquNdzFYfIbEXntcFbE=",
        version = "v1.9.2",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        sum = "h1:CXSaggrXdbHK9CF+8ywj8Amf7PBRmPCOJugH954Nnlo=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        sum = "h1:jFzHGLGAlb3ruxLB8MhbI6A8+AQX/2eW4qeyNZXNp2o=",
        version = "v1.0.6",
    )
    go_repository(
        name = "com_github_spiffe_go_spiffe_v2",
        importpath = "github.com/spiffe/go-spiffe/v2",
        sum = "h1:l+DolpxNWYgruGQVV0xsfeya3CsC7m8iBzDnMpsbLuo=",
        version = "v2.6.0",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:xuMeJ0Sdp5ZMRXx/aWO6RZxdr3beISkG5/G/aIRr3pY=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:Xv5erBjTwe/5IxqUQTdXv5kgmIvbHo3QQyRwhJsOfJA=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_tenntenn_golden",
        importpath = "github.com/tenntenn/golden",
        sum = "h1:3LPemp4rZUIYeh0MTgIn8nzbY8EKVu0Kyb7AOBM016g=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_thlib_go_timezone_local",
        importpath = "github.com/thlib/go-timezone-local",
        sum = "h1:BuzhfgfWQbX0dWzYzT1zsORLnHRv3bcRcsaUk0VmXA8=",
        version = "v0.0.0-20210907160436-ef149e42d28e",
    )
    go_repository(
        name = "com_github_tinylib_msgp",
        importpath = "github.com/tinylib/msgp",
        sum = "h1:FCXC1xanKO4I8plpHGH2P7koL/RzZs12l/+r7vakfm0=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_tklauser_go_sysconf",
        importpath = "github.com/tklauser/go-sysconf",
        sum = "h1:89WgdJhk5SNwJfu+GKyYveZ4IaJ7xAkecBo+KdJV0CM=",
        version = "v0.3.11",
    )
    go_repository(
        name = "com_github_tklauser_numcpus",
        importpath = "github.com/tklauser/numcpus",
        sum = "h1:kebhY2Qt+3U6RNK7UqpYNA+tJ23IBEGKkB7JQBfDYms=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_twitchyliquid64_golang_asm",
        importpath = "github.com/twitchyliquid64/golang-asm",
        sum = "h1:SU5vSMR7hnwNxj24w34ZyCi/FmDZTkS4MhqMhdFk5YI=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_ugorji_go_codec",
        importpath = "github.com/ugorji/go/codec",
        sum = "h1:Qd2W2sQawAfG8XSvzwhBeoGq71zXOC/Q1E9y/wUcsUA=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_vmware_labs_yaml_jsonpath",
        importpath = "github.com/vmware-labs/yaml-jsonpath",
        sum = "h1:/5QKeCBGdsInyDCyVNLbXyilb61MXGi9NP674f9Hobk=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_wk8_go_ordered_map_v2",
        importpath = "github.com/wk8/go-ordered-map/v2",
        sum = "h1:dLuIF2kX9c+KknGJUdJi1Il1SDiTSK158/BB9kdgAew=",
        version = "v2.1.9-0.20240815153524-6ea36470d1bd",
    )
    go_repository(
        name = "com_github_xanzy_go_gitlab",
        importpath = "github.com/xanzy/go-gitlab",
        sum = "h1:rWtwKTgEnXyNUGrOArN7yyc3THRkpYcKXIXia9abywQ=",
        version = "v0.15.0",
    )
    go_repository(
        name = "com_github_xdg_go_pbkdf2",
        importpath = "github.com/xdg-go/pbkdf2",
        sum = "h1:Su7DPu48wXMwC3bs7MCNG+z4FhcyEuz5dlvchbq0B0c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xdg_go_scram",
        importpath = "github.com/xdg-go/scram",
        sum = "h1:VOMT+81stJgXW3CpHyqHN3AXDYIMsx56mEFrB37Mb/E=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_xdg_go_stringprep",
        importpath = "github.com/xdg-go/stringprep",
        sum = "h1:kdwGpVNwPFtjs98xCGkHjQtGKh86rDcRZN17QEMCOIs=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonpointer",
        importpath = "github.com/xeipuuv/gojsonpointer",
        sum = "h1:zGWFAtiMcyryUHoUjUJX0/lt1H2+i2Ka2n+D3DImSNo=",
        version = "v0.0.0-20190905194746-02993c407bfb",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonreference",
        importpath = "github.com/xeipuuv/gojsonreference",
        sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
        version = "v0.0.0-20180127040603-bd5ef7bd5415",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonschema",
        importpath = "github.com/xeipuuv/gojsonschema",
        sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration_v2",
        importpath = "github.com/xhit/go-str2duration/v2",
        sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_xlab_treeprint",
        importpath = "github.com/xlab/treeprint",
        sum = "h1:HzHnuAF1plUN2zGlAFHbSQP2qJ0ZAD3XF5XD7OesXRQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xo_dburl",
        importpath = "github.com/xo/dburl",
        sum = "h1:NwFghJfjaUW7tp+WE5mTLQQCfgseRsvgXjlSvk7x4t4=",
        version = "v0.23.8",
    )
    go_repository(
        name = "com_github_xo_terminfo",
        importpath = "github.com/xo/terminfo",
        sum = "h1:JVG44RsyaB9T2KIHavMF/ppJZNG9ZpyihvCd0w101no=",
        version = "v0.0.0-20220910002029-abceb7e1c41e",
    )
    go_repository(
        name = "com_github_youmark_pkcs8",
        importpath = "github.com/youmark/pkcs8",
        sum = "h1:splanxYIlg+5LfHAM6xpdFEAYOk8iySO56hMFq6uLyA=",
        version = "v0.0.0-20181117223130-1be2e3e5546d",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:iERMLn0/QJeHFhxSt3p6PeN9mGnvIKSpG9YYorDMnic=",
        version = "v1.7.8",
    )
    go_repository(
        name = "com_github_yuin_goldmark_emoji",
        importpath = "github.com/yuin/goldmark-emoji",
        sum = "h1:EMVWyCGPlXJfUXBXpuMu+ii3TIaxbVBnEX9uaDC4cIk=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_yusufpapurcu_wmi",
        importpath = "github.com/yusufpapurcu/wmi",
        sum = "h1:KBNDSne4vP5mbSWnJbO+51IMOXJB67QiYCSBrubbPRg=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_zeebo_assert",
        importpath = "github.com/zeebo/assert",
        sum = "h1:g7C04CbJuIDKNPFHmsk4hwZDO5O+kntRxzaUoNXj+IQ=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_zeebo_errs",
        importpath = "github.com/zeebo/errs",
        sum = "h1:XNdoD/RRMKP7HD0UhJnIzUy74ISdGGxURlYG8HSWSfM=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_zeebo_xxh3",
        importpath = "github.com/zeebo/xxh3",
        sum = "h1:xZmwmqxHZA8AI603jOQ0tMqmBr9lPeFwGg6d+xy9DC0=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_gitlab_nyarla_go_crypt",
        importpath = "gitlab.com/nyarla/go-crypt",
        sum = "h1:7gd+rd8P3bqcn/96gOZa3F5dpJr/vEiDQYlNb/y2uNs=",
        version = "v0.0.0-20160106005555-d9a5dc2b789b",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:waZiuajrI28iAf40cWgycWNgaXPO06dupuS+sgibK6c=",
        version = "v0.121.6",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:Sc9ZjxFBEM/PoAxNlUwVGDcv8DYyjLYWDxHlzPG0q5I=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:2LnncRqfYB8NEdh9+FeYxAt9POTW/0zVboktnRlO11w=",
        version = "v1.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:SB4gurI6YSuZJCfAzjuIxxgbtJtczJFOTmCM7fLYRGg=",
        version = "v1.99.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:oYiABctb2wxUs1khR6vpC/T4P1EN27yvKI6Fxqzuv8E=",
        version = "v0.29.0",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:ehKUTy+QFsb3n07fEi18S2dpDDjCV4UlRyrbwfZV3Zk=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:S6s2zojwMymx0fyZYKm0eK1TdDxrriIBAlNVvRAOzug=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:TgdjAoGoRY81DEc2LYsYvi/OqCFImMzAk/TVKiSRsQw=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_apikeys",
        importpath = "cloud.google.com/go/apikeys",
        sum = "h1:B9CdHFZTFjVti89tmyXXrO+7vSNo2jvZuHG8zD5trdQ=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:IxGz6j5xv0nTJX285wu95Vn6KEi2CeV9vbyRgCSEAoU=",
        version = "v1.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:BbpzLwaIXVPorrrzTH+ni7P5mLemmPPfSZ7o39k7zQc=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:A20kj2S2HO9vlyBVyVFHPxArjxkXvLP5LjcdE7NhaPc=",
        version = "v1.17.1",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:i55wWC/EwVdHMyJgRfbLp/L6ez4nQuOpZwSxkuqN9ek=",
        version = "v1.21.1",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:ip/shfJYx6lrHBWYADjrrrubcm7uZzy50TTF5tPG7ek=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_auth",
        importpath = "cloud.google.com/go/auth",
        sum = "h1:mFWNQ2FEVWAliEQWpAdH80omXFokmrnbDhUS9cBywsI=",
        version = "v0.16.5",
    )
    go_repository(
        name = "com_google_cloud_go_auth_oauth2adapt",
        importpath = "cloud.google.com/go/auth/oauth2adapt",
        sum = "h1:keo8NaayQZ6wimpNSmW5OPc283g65QNIiLpZnkHRbnc=",
        version = "v0.2.8",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:ZLj48Ur2Qcso4M3bgOtjsOmeV5Ee92N14wuOc8OW+L0=",
        version = "v1.14.7",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:9bdGlpY1LgLONQjFsDwrkjLzdPTlROpfU+GhA97YpOk=",
        version = "v1.3.6",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:gWQdvdPplptpvrkqF6ibtxZkOsYKLTFbxYawHa/TvCg=",
        version = "v1.12.2",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:4FcR+4QmcNGkhVij6TrYS4AQVNLBo7PBXKxNrKzpclQ=",
        version = "v1.1.6",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:rZvHnjSUs5sHK3F9awiuFk2PeOaB8suqNuim21GbaTc=",
        version = "v1.69.0",
    )
    go_repository(
        name = "com_google_cloud_go_bigtable",
        importpath = "cloud.google.com/go/bigtable",
        sum = "h1:L/PnUXRtAzFfa7qMULJHt4cXa/O2dqPJEkzYNGA4hfo=",
        version = "v1.38.0",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:pqM5/c9UGydB9H90IPCxSvfCNLUPazAOSMsZkz5q5P4=",
        version = "v1.20.4",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:T0zYEroXT+y0O/x/yZd5SwQdFv4UbUINjvJyJKzDm0Q=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:+ZPglfDurCcsv4azizDFpBucD1IkRjWjbnU7zceyjfY=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:EeUa6SnD3+EL9B06G6N9Ud5/p/NtT6PC7lv5kmaUiHs=",
        version = "v1.20.0",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:FnbWf1FjtyrPLVqTb5wYX8eMgrbC5w0BPhpud9fAIu8=",
        version = "v1.22.3",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:IWJbQBEECTaNanDRN1XdR7FU53MJ1nylTl3s9T3MuyI=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:Fwan19UiNoFD+3KY0MnNHE5DyixOxNzS1mZ4ChOdpy0=",
        version = "v1.13.6",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:6gRrxftrqe5llEyTvwPGGEqTnetXOrlLhPPyU4oTd34=",
        version = "v1.43.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:HxMRIbao8w17ZX6wBnjhcDkW6lTFpgcaobyVfZWqRLA=",
        version = "v0.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:lenyU3uzHwKDveCwmpfNxHYvLS3uEBWdn+O7+rSxy+Q=",
        version = "v1.17.3",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        importpath = "cloud.google.com/go/container",
        sum = "h1:JEHeW535svvNwJrjrlQ/cdjd15LCWrPKnHsulrufd3A=",
        version = "v1.44.0",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:1SoHlNqL3XrhqcoozB+3eoHif2sRUFtp/JeASQTtGKo=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:eFgygb3DTufTWWUB8ARk+dSuXz+aefNJXTlkWlQcWwE=",
        version = "v1.26.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:AdhB4cAkMOC9NtrHJxpKOVvO/VqBLaIyk0tEEhbGjYM=",
        version = "v0.11.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:0eCPTPUC/RZ863aVfXTJLkg0tEpdpn62VD6ywSmmzxM=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:GZ6J+CR8CEeWAj8luRCtr8GvImSQRkArIIqGiZOnzBA=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:VOZ5U+78ttnhNCEID7qdeogqZQzK5N+LPHIQ9Q3YDsc=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:nu8/KrLR5v62L1lApGNgm61Oq+xaa2bS9rgc1csjqE0=",
        version = "v1.26.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc",
        importpath = "cloud.google.com/go/dataproc",
        sum = "h1:W47qHL3W4BPkAIbk4SWmIERwsWBaNnWm0P2sdx3YgGU=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:oiEM2efaJfiOClBOYcmW1K+tDP5pHdnYsPmfD55tiRw=",
        version = "v2.14.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:qTRAG/E3T63Xj1orefRlwupfwH9c9ERUAnWSRGp75so=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:NNpXoyEqIJmZFc0ACcwBEaXnmscUpcG4NkKnbCePmiM=",
        version = "v1.20.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:ZXT4aQ90rXBy6B4UjiEYDdfZULpAdP7TmAxa2ZBwf28=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:C0VqBhFyQFp6+xgPHZAD7LeRA4XGy5YLzGmPQ2NhlLk=",
        version = "v1.27.2",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:nW3vH/ysZWBdjQJ4rIh3PC5Do/Brz8KEp3OeDy9VW3U=",
        version = "v1.69.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:ThCQO8Qy5TAfFEJQjhq80u5c93UMdM2uqI3pUZVy7Do=",
        version = "v1.24.0",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:7fla8GcarupO15eatRTUveXCob6DOSW1Wa+1i63CM3Q=",
        version = "v1.37.0",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:TI+Aavwc31KD8huOquJz0ISchCq1zSEWc9M+JcPJyxc=",
        version = "v0.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:9tfGCicvrki927T+hGMB0yYmwIbRuZY6JR1/awrKiZ0=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:isaoPwWX8kbAOea4qahcmttoS79+gQhvKsfg5L5AgH8=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:ysHZ4gr4plW1CL1Ur/AucUUfh20hDjSFbfjxSK0q/sk=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:bZW7ZMM+XXNErg6rOZcgxUzAgz4vpReRDP3ZiGf7/sI=",
        version = "v1.15.5",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:LjoAyp9TvVNBns3sUUzPaNsQiGpR2BReGmTS3bUCuBE=",
        version = "v1.10.2",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:cuydCaLS7Vl2SatAeivXyhbhDEIR8BDmtn4egDhIn2s=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:vJgWlvxtJG6p/JrbXAkz83DbgwOyFhZZI1Y32vUddjY=",
        version = "v1.19.6",
    )
    go_repository(
        name = "com_google_cloud_go_gaming",
        importpath = "cloud.google.com/go/gaming",
        sum = "h1:7vEhFnZmd931Mo7sZ6pJy7uQPDxF7m7v8xtBheG08tc=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:eBqOt61yEChvj7I/GDPBbdCCRdUPudD1qrQYfYWV3Ok=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:67/rnPmF/I1Wmf7jWyKH+z4OWjU8ZUI0Vmzxvmzf3KY=",
        version = "v0.12.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:9iogrmNNa+drDPf/zkLH/6KGgUf7FuuyokmithoGwMQ=",
        version = "v0.15.6",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:334aZmOzIt3LVBpguCof8IHaLaftcZlx+L0TGBukYkY=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_google_cloud_go_grafeas",
        importpath = "cloud.google.com/go/grafeas",
        sum = "h1:lBjwKmhpiqOAFaE0xdqF8CqO74a99s8tUT5mCkBBxPs=",
        version = "v0.3.15",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:sk0SxpCGIA7tIO//XdiiG29f2vrF6Pq/dsxxyBGiRBY=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:qgFRAGEmd8z6dJ/qyEchAuL9jpswyODjA2lS+w234g8=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:VIioCrYsyWiRGx7Y8RDNylpI6d4t1Qx5ZgSLUVmWWPo=",
        version = "v1.11.2",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:uKGuaWozDcjg3wyf54Gd7tCH2YK8BFeH9qo1xBNiPKE=",
        version = "v1.5.6",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:A3AhugnIViAZkC3/lHAQDaXBIk2ZOPBZS0XQCyZsjjc=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:dBRIj7+GDeeEvatJeTB19oYZNV0aj6wEqSIT/7gLqtk=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        importpath = "cloud.google.com/go/language",
        sum = "h1:BVJ/POtlnJ55LElvnQY19UOxpMVtHoHHkFJW2uHJsVU=",
        version = "v1.14.5",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:Vu7XF4s5KJ8+mSLIL4eaQM6JTyWXvSB54oqC+CUZH20=",
        version = "v0.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:7j0HgAp0B94o1YRDqiqm26w4q1rDMH7XNRU34lJXHYc=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:IGtfDWHhQCgCjwQjV9iiLnUta9LBCo8R9QmAFsS/PrE=",
        version = "v0.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:zrZVWXZJlmHnfpyCrTQIbDBGUBHrcOOvrsjMjoXRxrk=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:Rgs6jvqYztt28cKTh9hzk/a2qq/83FrqR92nPM3bb10=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:SDGatA73TgZ8iCvILVXpk/1qhTK5DJyufUDEWgbmbV8=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:33IVqQEmFiITsBXwGHeTkUhWz0kLNKr90nV3e22uLPs=",
        version = "v1.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:dLm59AHHZCorveCylj7c2iWhkQsmMIeWTsV+tG/BXtY=",
        version = "v1.14.7",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:5OTsoJ1dXYIiMiuL+sYscLc9BumrL3CarVLL7dd7lHM=",
        version = "v1.24.2",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:dQxh+/LBr8heinPCwjlYHZAr4IU7qCusl5yMK9+9apU=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:xU5mNtugdHmeZfsJKDi+agkKC3joY2Mp750AAhqk9Qs=",
        version = "v1.20.0",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:6b6fcCG9BFNcmtNO+VuPE04vkZb5TKNX9+7ZhYMgstE=",
        version = "v0.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:nCfZwVihArMPP2atRoxRrXOXJ/aC9rAgpBQGCc2zpYw=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:jDvIuSxDsXI2P7l2sYXm6CoX1YBIIT6Khm5m0hq0/KQ=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:PnlZ/O4R/eiounpxUkhI9ZXRMWbG7vFqxc6L6sR+31k=",
        version = "v1.11.9",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:uQziDu3UKYk9ZwUgneZAW5aWxZFKgOXXsuVKFKh0z7Y=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:7ol0qfm+vzfgNtGcLVJevs7UW/7JR+MlmPI/LqoZrUM=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:BDKVcxo1OO4ZT+PbuFchZjnbrlUGfChilt6+pITY1VI=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:yl572bBQbPjflX250SOflN6gwO2uYoddN2uRp36fDTo=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:Z8+tO2z21MY1arBBuJjwrOjbw8fbZb13AZTHXdzkl2U=",
        version = "v1.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:R951ikhxIanXEijBCu0xnoUAOteS5m/Xplek0YvsNTE=",
        version = "v0.10.7",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:hnYpOIxVlgVD1Z8LN7est4DQZK3K6tvZNurZjIVjUe0=",
        version = "v1.50.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub_v2",
        importpath = "cloud.google.com/go/pubsub/v2",
        sum = "h1:0qS6mRJ41gD1lNmM/vdm6bR7DQu6coQcVwD+VPf0Bz0=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:jLQozsEVr+c6tOU13vDugtnaBSUy/PD5zK6mhm+uF1Y=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise",
        importpath = "cloud.google.com/go/recaptchaenterprise",
        sum = "h1:u6EznTGzIdsyOsvm+Xkw0aSuKFXQlyjGE9a4exk6iNQ=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:P4QMryKcWdi4LIe1Sx0b2ZOAQv5gVfdzPt2peXcN32Y=",
        version = "v2.20.4",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:slN7h23vswGccW8x3f+xUXCu9Yo18/GNkazH93LJbFk=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:cIsyRKGNw4LpCfY5c8CCQadhlp54jP4fHtP+d5Sy2xE=",
        version = "v1.13.5",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:JlHLceAOILEmbn+NIS7l+vmUKkFuobLToCWTxL7NGcQ=",
        version = "v1.18.2",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:LIa8kKE8HF71zm976oHMqpWFiaDHVw/H1YMO71lrGmo=",
        version = "v1.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:13HOFU7v4cEvIHXSAQbinF4wp2Baybbq7q9FMctg1Ek=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:2jXyUed1nbv4aL2drKF/HbQnN1TgEorWHp1EAxNNetg=",
        version = "v1.24.0",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        importpath = "cloud.google.com/go/run",
        sum = "h1:l4tpqhzJ75uOugXl2BQ15uEM5gLamVH5M70tBv70ZCU=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:zkMEJ0UbEJ3O7NwEUlKLIp6eXYv1L7wHjbxyxznajKM=",
        version = "v1.11.7",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:RtkCMgTpaBMbzozcRUGfZe46jb9a3qh5EdEtVRUATF8=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        importpath = "cloud.google.com/go/security",
        sum = "h1:GI6kufPbHFINq998M7x60rfV5MVo6yhD+uVrjNzueKw=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:UR8cUgXFYpWxKkKnUNy65hlrAzgwBZBVxZilJ50ESXU=",
        version = "v1.37.0",
    )
    go_repository(
        name = "com_google_cloud_go_servicecontrol",
        importpath = "cloud.google.com/go/servicecontrol",
        sum = "h1:d0uV7Qegtfaa7Z2ClDzr9HJmnbJW7jn0WhZ7wOX6hLE=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:pl/KUNvFzlXpxgnPgzQjyTQQcv5WsQ97zCHaPrLQlYA=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_servicemanagement",
        importpath = "cloud.google.com/go/servicemanagement",
        sum = "h1:fopAQI/IAzlxnVeiKn/8WiV6zKndjFkvi+gzu+NjywY=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_serviceusage",
        importpath = "cloud.google.com/go/serviceusage",
        sum = "h1:rXyq+0+RSIm3HFypctp7WoXxIA563rn206CfMWdqXX4=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:jLWyztGlNWBx55QXBM4HbWvfv7aiRjPzRKTUkZA8dXk=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:ShH4Y3YeDtmHa55dFiSS3YtQ0dmCuP0okfAoHp/d68w=",
        version = "v1.84.1",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:9AuiAxDTmh/aeREtw+/0e7aI27T5QN4fK5lhssc9MxA=",
        version = "v1.28.0",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:n6gy+yLnHn0hTwBFzNn8zJ1kqWfR91wzdM8hjRF4wP0=",
        version = "v1.56.1",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:uqKX3OgcYzR1W1YI943ZZ45id0RqA2eXXoCBSPstlbw=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:wDP+++O/P1cTJBMkYlSY46k0a6atSoyO+UkBGuU9+Ao=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:oWWFQp0yFl4EJOr3opDkKH9304wUsZjgPjrTDS6S1a8=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:S4Ptq+yFIPNLEzQ/OQwiIYDNzk5I2vYmhf0SmFQOmWo=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:2O2zjPzqPYAHrn3OKl029qlqG6W8ZdYaOWRyr8NgMT4=",
        version = "v1.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:QHcszWZvBLEZHM2WJ6IDg2BUTWzEPMiHhbJAd15yKGU=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        importpath = "cloud.google.com/go/video",
        sum = "h1:M/krJJDgUgZszzFgv7qHRltSVMafXgcuYyCETSH1oyA=",
        version = "v1.25.0",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:heq7jEO39sH5TycBh8TGFJ827XCxK0tIWatmBY/n0jI=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_vision",
        importpath = "cloud.google.com/go/vision",
        sum = "h1:/CsSTkbmO9HC8iQpxbK8ATms3OQaX3YQUeTMGCxlaK4=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:UJZ0H6UlOaYKgCn6lWG2iMAOJIsJZLnseEfzBR8yIqQ=",
        version = "v2.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:68hOQDhs1DOITrCrhritrwr8xy6s8QMdwDyMzMiFleU=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:OsGd1SB91y9fDuzdzFngMv4UcT4cqmRxjsCsS4Xmcu8=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:RYtUB9rQEijX9Tc6lQcGst58ZOzPgaYTkz6+2pyPQTM=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:yZKNB7zRxOMriLrhP5WDE+BjxXVl0wJHHZSdaYzbdVU=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:cIPKJKZA3l7D8DfL4nxce8HGOWXBw3WAUBF0ymOW9GQ=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:phBz5TOAES0YGogxZ6Q7ISSudaf618lRhE3euzBpE9U=",
        version = "v1.14.2",
    )
    go_repository(
        name = "com_google_firebase_go_v4",
        importpath = "firebase.google.com/go/v4",
        sum = "h1:S+g0P72oDGqOaG4wlLErX3zQmU9plVdu7j+Bc3R1qFw=",
        version = "v4.18.0",
    )
    go_repository(
        name = "com_lukechampine_uint128",
        importpath = "lukechampine.com/uint128",
        sum = "h1:cDdUVfRwDUDovz610ABgFD17nXD4/uDgVHl2sC3+sbo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_shuralyov_dmitri_gpu_mtl",
        importpath = "dmitri.shuralyov.com/gpu/mtl",
        sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
        version = "v0.0.0-20190408044501-666a987793e9",
    )
    go_repository(
        name = "dev_cel_expr",
        importpath = "cel.dev/expr",
        sum = "h1:56OvJKSH3hDGL0ml5uSxZmz3/3Pq4tJ+fb1unVLAFcY=",
        version = "v0.24.0",
    )
    go_repository(
        name = "ht_sr_git_sbinet_gg",
        importpath = "git.sr.ht/~sbinet/gg",
        sum = "h1:RIzgkizAk+9r7uPzf/VfbJHBMKUr0F5hRFxTUGMnt38=",
        version = "v0.6.0",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_errgo_v2",
        importpath = "gopkg.in/errgo.v2",
        sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version = "v2.1.0",
    )
    go_repository(
        name = "in_gopkg_h2non_gock_v1",
        importpath = "gopkg.in/h2non/gock.v1",
        sum = "h1:jBbHXgGBK/AoPVfJh5x4r/WxIrElvbLel8TCZkkZJoY=",
        version = "v1.1.2",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        importpath = "gopkg.in/inf.v0",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        version = "v0.9.1",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=",
        version = "v1.67.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "in_yaml_go_yaml_v3",
        importpath = "go.yaml.in/yaml/v3",
        sum = "h1:tfq32ie2Jv2UxXFdLJdh3jXuOzWiL1fo0bu/FbuKpbc=",
        version = "v3.0.4",
    )
    go_repository(
        name = "in_yaml_go_yaml_v4",
        importpath = "go.yaml.in/yaml/v4",
        sum = "h1:4J1+yLKUIPGexM/Si+9d3pij4hdc7aGO04NhrElqXbY=",
        version = "v4.0.0-rc.1",
    )
    go_repository(
        name = "io_etcd_go_etcd_api_v3",
        importpath = "go.etcd.io/etcd/api/v3",
        sum = "h1:sbcmosSVesNrWOJ58ZQFitHMdncusIifYcrBfwrlJSY=",
        version = "v3.5.7",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_pkg_v3",
        importpath = "go.etcd.io/etcd/client/pkg/v3",
        sum = "h1:y3kf5Gbp4e4q7egZdn5T7W9TSHUvkClN6u+Rq9mEOmg=",
        version = "v3.5.7",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v3",
        importpath = "go.etcd.io/etcd/client/v3",
        sum = "h1:u/OhpiuCgYY8awOHlhIhmGIGpxfBU/GZBUP3m/3/Iz4=",
        version = "v3.5.7",
    )
    go_repository(
        name = "io_filippo_edwards25519",
        importpath = "filippo.io/edwards25519",
        sum = "h1:FNf4tywRC1HmFuKW5xopWpigGjJKiJSV0Cqo0cJWDaA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "io_gorm_driver_mysql",
        importpath = "gorm.io/driver/mysql",
        sum = "h1:eNbLmNTpPpTOVZi8MMxCi2aaIm0ZpInbORNXDwyLGvg=",
        version = "v1.6.0",
    )
    go_repository(
        name = "io_gorm_driver_postgres",
        importpath = "gorm.io/driver/postgres",
        sum = "h1:2dxzU8xJ+ivvqTRph34QX+WrRaJlmfyPqXmoGVjMBa4=",
        version = "v1.6.0",
    )
    go_repository(
        name = "io_gorm_driver_sqlserver",
        importpath = "gorm.io/driver/sqlserver",
        sum = "h1:XWISFsu2I2pqd1KJhhTZNJMx1jNQ+zVL/Q8ovDcUjtY=",
        version = "v1.6.1",
    )
    go_repository(
        name = "io_gorm_gorm",
        importpath = "gorm.io/gorm",
        sum = "h1:lSHg33jJTBxs2mgJRfRZeLDG+WZaHYCk3Wtfl6Ngzo4=",
        version = "v1.30.1",
    )
    go_repository(
        name = "io_gorm_plugin_dbresolver",
        importpath = "gorm.io/plugin/dbresolver",
        sum = "h1:F4b85TenghUeITqe3+epPSUtHH7RIk3fXr5l83DF8Pc=",
        version = "v1.6.2",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opencensus_go_contrib_exporter_stackdriver",
        importpath = "contrib.go.opencensus.io/exporter/stackdriver",
        sum = "h1:xRc46S76eyn4ZF3jWX8I+aUSKVLw5EQ1aDvHwfV5W1o=",
        version = "v0.13.15-0.20230702191903-2de6d2748484",
    )
    go_repository(
        name = "io_opentelemetry_go_auto_sdk",
        importpath = "go.opentelemetry.io/auto/sdk",
        sum = "h1:cH53jehLUN6UFLY71z+NDOiNJqDdPRaXzTel0sJySYA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_detectors_gcp",
        importpath = "go.opentelemetry.io/contrib/detectors/gcp",
        sum = "h1:B+WbN9RPsvobe6q4vP6KgM8/9plR/HNjgGBrfcOlweA=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_github_com_gin_gonic_gin_otelgin",
        importpath = "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin",
        sum = "h1:fZNpsQuTwFFSGC96aJexNOBrCD7PjD9Tm/HyHtXhmnk=",
        version = "v0.62.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:rbRJ8BBoVMsQShESYZ0FkvcITu8X8QNwJogcLUmDNNw=",
        version = "v0.62.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
        sum = "h1:Hf9xI/XLML9ElpiHVDNwvqI0hIFlzV8dgIr35kV1kRU=",
        version = "v0.62.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_propagators_b3",
        importpath = "go.opentelemetry.io/contrib/propagators/b3",
        sum = "h1:0aGKdIuVhy5l4GClAjl72ntkZJhijf2wg1S7b5oLoYA=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:9zhNfelUvx0KBfu/gb+ZgeAfAgtWrfHJZcAqFC228wQ=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
        sum = "h1:Ahq7pZmv87yiyn3jeFz/LekZmPLLdKejuO3NcK9MssM=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracehttp",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp",
        sum = "h1:bDMKF3RUSxshZ5OjOTi8rsHGaPKsAt76FaqgvIUySLc=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_prometheus",
        importpath = "go.opentelemetry.io/otel/exporters/prometheus",
        sum = "h1:AHh/lAP1BHrY5gBwk8ncc25FXWm/gmmY3BX258z5nuk=",
        version = "v0.57.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_stdout_stdoutmetric",
        importpath = "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric",
        sum = "h1:rixTyDGXFxRy1xzhKrotaHy3/KXdPhlWARrCgK+eqUY=",
        version = "v1.36.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_stdout_stdouttrace",
        importpath = "go.opentelemetry.io/otel/exporters/stdout/stdouttrace",
        sum = "h1:SNhVp/9q4Go/XHBkQ1/d5u9P/U+L1yaGPoi0x+mStaI=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:mvwbQS5m0tbmqML4NqK+e3aDiO02vsf/WgbsdpcPoZE=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:ItB0QUqnjesGRvNcmAcU0LyvkVyGJ2xftD29bWdDvKI=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_metric",
        importpath = "go.opentelemetry.io/otel/sdk/metric",
        sum = "h1:90lI228XrB9jCMuSdA0673aubgRobVZFhbjxHHspCPc=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:HLdcFNbRQBE2imdSEgm/kwqmQj1Or1l/7bW6mxVK7z4=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_proto_otlp",
        importpath = "go.opentelemetry.io/proto/otlp",
        sum = "h1:gTOMpGDb0WTBOP8JaO72iL3auEZhVmAQg4ipjOVAtj4=",
        version = "v1.7.1",
    )
    go_repository(
        name = "io_rsc_binaryregexp",
        importpath = "rsc.io/binaryregexp",
        sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "io_rsc_pdf",
        importpath = "rsc.io/pdf",
        sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
        version = "v0.1.1",
    )
    go_repository(
        name = "io_rsc_quote_v3",
        importpath = "rsc.io/quote/v3",
        sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
        version = "v3.1.0",
    )
    go_repository(
        name = "io_rsc_sampler",
        importpath = "rsc.io/sampler",
        sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_codeberg_go_fonts_liberation",
        importpath = "codeberg.org/go-fonts/liberation",
        sum = "h1:SsKoMO1v1OZmzkG2DY+7ZkCL9U+rrWI09niOLfQ5Bo0=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_codeberg_go_latex_latex",
        importpath = "codeberg.org/go-latex/latex",
        sum = "h1:hoGO86rIbWVyjtlDLzCqZPjNykpWQ9YuTZqAzPcfL3c=",
        version = "v0.1.0",
    )
    go_repository(
        name = "org_codeberg_go_pdf_fpdf",
        importpath = "codeberg.org/go-pdf/fpdf",
        sum = "h1:u+w669foDDx5Ds43mpiiayp40Ov6sZalgcPMDBcZRd4=",
        version = "v0.10.0",
    )
    go_repository(
        name = "org_gioui",
        importpath = "gioui.org",
        sum = "h1:K72hopUosKG3ntOPNG4OzzbuhxGuVf06fa2la1/H/Ho=",
        version = "v0.0.0-20210308172011-57750fc8a0a6",
    )
    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        sum = "h1:hUotakSkcwGdYUqzCRc5yGYsg4wXxpkKlW5ryVqvC1Y=",
        version = "v0.248.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_appengine_v2",
        importpath = "google.golang.org/appengine/v2",
        sum = "h1:LvPZLGuchSBslPBp+LAhihBeGSiRh1myRoYK4NtuBIw=",
        version = "v2.0.6",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:ZERoum3uuqL0PRSc6SXielu26FN96T4BUGaaW0oL+c8=",
        version = "v0.0.0-20250818200422-3122310a409c",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:AtEkQdl5b6zsybXcbz00j1LwNodDuH6hVifIaNqk7NQ=",
        version = "v0.0.0-20250818200422-3122310a409c",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:CMCT63H4Rl6uNNT10m3hkjCR3JgAv4E9ZuVTeO+Sz98=",
        version = "v0.0.0-20250818200422-3122310a409c",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:qXWI/sQtv5UKboZ/zUk7h+mrf/lXORyI+n9DKDAusdg=",
        version = "v0.0.0-20250818200422-3122310a409c",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:+TW+dqTd2Biwe6KKfhE5JpiYIBWq865PhKGSXiivqt4=",
        version = "v1.75.0",
    )
    go_repository(
        name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
        importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
        sum = "h1:M1YKkFIboKNieVO5DLUEVzQfGwJD30Nv2jfUgzb5UcE=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_golang_google_grpc_examples",
        importpath = "google.golang.org/grpc/examples",
        sum = "h1:ExN12ndbJ608cboPYflpTny6mXSzPrDLh0iTaVrRrds=",
        version = "v0.0.0-20250407062114-b368379ef8f6",
    )
    go_repository(
        name = "org_golang_google_grpc_gcp_observability",
        importpath = "google.golang.org/grpc/gcp/observability",
        sum = "h1:2IQ7szW1gobfZaS/sDSAu2uxO0V/aTryMZvlcyqKqQA=",
        version = "v1.0.1",
    )
    go_repository(
        name = "org_golang_google_grpc_security_advancedtls",
        importpath = "google.golang.org/grpc/security/advancedtls",
        sum = "h1:/KQ7VP/1bs53/aopk9QhuPyFAp9Dm9Ejix3lzYkCrDA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_golang_google_grpc_stats_opencensus",
        importpath = "google.golang.org/grpc/stats/opencensus",
        sum = "h1:evSYcRZaSToQp+borzWE52+03joezZeXcKJvZDfkUJA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:xHScyCOEuuwZEc6UtSOvPbAT4zRh0xcNRYekJwfqyMc=",
        version = "v1.36.8",
    )
    go_repository(
        name = "org_golang_x_arch",
        importpath = "golang.org/x/arch",
        sum = "h1:dx1zTU0MAE98U+TQ8BLl7XsJbgze2WnNKF/8tGp/Q6c=",
        version = "v0.20.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:WKYxWedPGCTVVl5+WHSSrOBT0O8lx32+zxmHxijgXp4=",
        version = "v0.41.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:DXr+pvt3nC887026GRP39Ej11UATqWDmWuS99x26cD0=",
        version = "v0.0.0-20250819193227-8b4c13bb791b",
    )
    go_repository(
        name = "org_golang_x_image",
        importpath = "golang.org/x/image",
        sum = "h1:Y6uW6rH1y5y/LK1J8BPWZtr6yZ7hrsy6hFrXjgsc2fQ=",
        version = "v0.25.0",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
        version = "v0.0.0-20210508222113-6edffad5e616",
    )
    go_repository(
        name = "org_golang_x_mobile",
        importpath = "golang.org/x/mobile",
        sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
        version = "v0.0.0-20190719004257-d2bd2a29d028",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:kb+q2PyFnEADO2IEF935ehFUXlWiNjJWtRNgBLSfbxQ=",
        version = "v0.27.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:lat02VYK2j4aLzMzecihNvTlJNQUq316m2Mr9rnM6YE=",
        version = "v0.43.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:dnDm7JmhM45NNpd8FDDeLhK6FwqbOf4MLCM9zb1BOHI=",
        version = "v0.30.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:ycBJEhp9p4vXvUZNszeOq0kGTPghopOL8q0fq3vstxw=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:vz1N37gP5bs89s7He8XuIYXpyY0+QlsKmzipCbUtyxI=",
        version = "v0.35.0",
    )
    go_repository(
        name = "org_golang_x_telemetry",
        importpath = "golang.org/x/telemetry",
        sum = "h1:3doPGa+Gg4snce233aCWnbZVFsyFMo/dR40KK/6skyE=",
        version = "v0.0.0-20250807160809-1a19826ec488",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:O/2T7POpk0ZZ7MAzMeWFSg6S5IpWd/RXDlM9hgM3DR4=",
        version = "v0.34.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:rhazDwis8INMIwQ4tpjLDzUhx6RlXqZNPEM0huQojng=",
        version = "v0.28.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:ScB/8o8olJvc+CQPWrK3fPZNfh7qgwCrY0zJmoEQLSE=",
        version = "v0.12.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:kWS0uv/zsvHEle1LbV5LE8QujrxB3wfQyxHfhOk0Qkg=",
        version = "v0.36.0",
    )
    go_repository(
        name = "org_golang_x_tools_go_expect",
        importpath = "golang.org/x/tools/go/expect",
        sum = "h1:jpBZDwmgPhXsKZC6WhL20P4b/wmnpsEAGHaNy0n/rJM=",
        version = "v0.1.1-deprecated",
    )
    go_repository(
        name = "org_golang_x_tools_go_packages_packagestest",
        importpath = "golang.org/x/tools/go/packages/packagestest",
        sum = "h1:1h2MnaIAIXISqTFKdENegdpAgUXz6NrPEsbIeWaBRvM=",
        version = "v0.1.1-deprecated",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:noIWHXmPHxILtqtCOPIhSt0ABwskkZKjD3bXGnZGpNY=",
        version = "v0.0.0-20240903120638-7835f813f4da",
    )
    go_repository(
        name = "org_gonum_v1_gonum",
        importpath = "gonum.org/v1/gonum",
        sum = "h1:5+ul4Swaf3ESvrOnidPp4GZbzf0mxVQpDCYUQE7OJfk=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_gonum_v1_netlib",
        importpath = "gonum.org/v1/netlib",
        sum = "h1:OE9mWmgKkjJyEmDAAtGMPjXu+YNeGvK9VTSHY6+Qihc=",
        version = "v0.0.0-20190313105609-8cb42192e0e0",
    )
    go_repository(
        name = "org_gonum_v1_plot",
        importpath = "gonum.org/v1/plot",
        sum = "h1:Tlfh/jBk2tqjLZ4/P8ZIwGrLEWQSPDLRm/SNWKNXiGI=",
        version = "v0.15.2",
    )
    go_repository(
        name = "org_modernc_b",
        importpath = "modernc.org/b",
        sum = "h1:vpvqeyp17ddcQWF29Czawql4lDdABCDRbXRAS4+aF2o=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_cc_v3",
        importpath = "modernc.org/cc/v3",
        sum = "h1:QoR1Sn3YWlmA1T4vLaKZfawdVtSiGx8H+cEojbC7v1Q=",
        version = "v3.41.0",
    )
    go_repository(
        name = "org_modernc_cc_v4",
        importpath = "modernc.org/cc/v4",
        sum = "h1:yEN8dzrkRFnn4PUUKXLYIqVf2PJYAEjMTFjO3BDGc3I=",
        version = "v4.26.3",
    )
    go_repository(
        name = "org_modernc_ccgo_v3",
        importpath = "modernc.org/ccgo/v3",
        sum = "h1:o3OmOqx4/OFnl4Vm3G8Bgmqxnvxnh0nbxeT5p/dWChA=",
        version = "v3.17.0",
    )
    go_repository(
        name = "org_modernc_ccgo_v4",
        importpath = "modernc.org/ccgo/v4",
        sum = "h1:rjznn6WWehKq7dG4JtLRKxb52Ecv8OUGah8+Z/SfpNU=",
        version = "v4.28.0",
    )
    go_repository(
        name = "org_modernc_ccorpus",
        importpath = "modernc.org/ccorpus",
        sum = "h1:J16RXiiqiCgua6+ZvQot4yUuUy8zxgqbqEEUuGPlISk=",
        version = "v1.11.6",
    )
    go_repository(
        name = "org_modernc_ccorpus2",
        importpath = "modernc.org/ccorpus2",
        sum = "h1:Ui+4tc58mf/W+2arcYCJR903y3zl3ecsI7Fpaaqozyw=",
        version = "v1.5.2",
    )
    go_repository(
        name = "org_modernc_db",
        importpath = "modernc.org/db",
        sum = "h1:2c6NdCfaLnshSvY7OU09cyAY0gYXUZj4lmg5ItHyucg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_file",
        importpath = "modernc.org/file",
        sum = "h1:9/PdvjVxd5+LcWUQIfapAWRGOkDLK90rloa8s/au06A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_fileutil",
        importpath = "modernc.org/fileutil",
        sum = "h1:rJAXTP6ilMW/1+kzDiqmBlHLWszheUFXIyGQIAvjJpY=",
        version = "v1.3.15",
    )
    go_repository(
        name = "org_modernc_gc_v2",
        importpath = "modernc.org/gc/v2",
        sum = "h1:nyqdV8q46KvTpZlsw66kWqwXRHdjIlJOhG6kxiV/9xI=",
        version = "v2.6.5",
    )
    go_repository(
        name = "org_modernc_goabi0",
        importpath = "modernc.org/goabi0",
        sum = "h1:HvEowk7LxcPd0eq6mVOAEMai46V+i7Jrj13t4AzuNks=",
        version = "v0.2.0",
    )
    go_repository(
        name = "org_modernc_golex",
        importpath = "modernc.org/golex",
        sum = "h1:wWpDlbK8ejRfSyi0frMyhilD3JBvtcx2AdGDnU+JtsE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_httpfs",
        importpath = "modernc.org/httpfs",
        sum = "h1:AAgIpFZRXuYnkjftxTAZwMIiwEqAfk8aVB2/oA6nAeM=",
        version = "v1.0.6",
    )
    go_repository(
        name = "org_modernc_internal",
        importpath = "modernc.org/internal",
        sum = "h1:XMDsFDcBDsibbBnHB2xzljZ+B1yrOVLEFkKL2u15Glw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_lex",
        importpath = "modernc.org/lex",
        sum = "h1:prSCNTLw1R4rn7M/RzwsuMtAuOytfyR3cnyM07P+Pas=",
        version = "v1.1.1",
    )
    go_repository(
        name = "org_modernc_lexer",
        importpath = "modernc.org/lexer",
        sum = "h1:hU7xVbZsqwPphyzChc7nMSGrsuaD2PDNOmzrzkS5AlE=",
        version = "v1.0.4",
    )
    go_repository(
        name = "org_modernc_libc",
        importpath = "modernc.org/libc",
        sum = "h1:rjhZ8OSCybKWxS1CJr0hikpEi6Vg+944Ouyrd+bQsoY=",
        version = "v1.66.7",
    )
    go_repository(
        name = "org_modernc_lldb",
        importpath = "modernc.org/lldb",
        sum = "h1:6vjDJxQEfhlOLwl4bhpwIz00uyFK4EmSYcbwqwbynsc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_mathutil",
        importpath = "modernc.org/mathutil",
        sum = "h1:GCZVGXdaN8gTqB1Mf/usp1Y/hSqgI2vAGGP4jZMCxOU=",
        version = "v1.7.1",
    )
    go_repository(
        name = "org_modernc_memory",
        importpath = "modernc.org/memory",
        sum = "h1:o4QC8aMQzmcwCK3t3Ux/ZHmwFPzE6hf2Y5LbkRs+hbI=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_modernc_opt",
        importpath = "modernc.org/opt",
        sum = "h1:2kNGMRiUjrp4LcaPuLY2PzUfqM/w9N23quVwhKt5Qm8=",
        version = "v0.1.4",
    )
    go_repository(
        name = "org_modernc_ql",
        importpath = "modernc.org/ql",
        sum = "h1:bIQ/trWNVjQPlinI6jdOQsi195SIturGo3mp5hsDqVU=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_scannertest",
        importpath = "modernc.org/scannertest",
        sum = "h1:JPtfxcVdbRvzmRf2YUvsDibJsQRw8vKA/3jb31y7cy0=",
        version = "v1.0.2",
    )
    go_repository(
        name = "org_modernc_sortutil",
        importpath = "modernc.org/sortutil",
        sum = "h1:+xyoGf15mM3NMlPDnFqrteY07klSFxLElE2PVuWIJ7w=",
        version = "v1.2.1",
    )
    go_repository(
        name = "org_modernc_sqlite",
        importpath = "modernc.org/sqlite",
        sum = "h1:Aclu7+tgjgcQVShZqim41Bbw9Cho0y/7WzYptXqkEek=",
        version = "v1.38.2",
    )
    go_repository(
        name = "org_modernc_strutil",
        importpath = "modernc.org/strutil",
        sum = "h1:UneZBkQA+DX2Rp35KcM69cSsNES9ly8mQWD71HKlOA0=",
        version = "v1.2.1",
    )
    go_repository(
        name = "org_modernc_tcl",
        importpath = "modernc.org/tcl",
        sum = "h1:npxzTwFTZYM8ghWicVIX1cRWzj7Nd8i6AqqX2p+IYao=",
        version = "v1.13.1",
    )
    go_repository(
        name = "org_modernc_token",
        importpath = "modernc.org/token",
        sum = "h1:Xl7Ap9dKaEs5kLoOQeQmPWevfnk/DM5qcLcYlA8ys6Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_modernc_z",
        importpath = "modernc.org/z",
        sum = "h1:RTNHdsrOpeoSeOF4FbzTo8gBYByaJ5xT7NgZ9ZqRiJM=",
        version = "v1.5.1",
    )
    go_repository(
        name = "org_modernc_zappy",
        importpath = "modernc.org/zappy",
        sum = "h1:dPVaP+3ueIUv4guk8PuZ2wiUGcJ1WUVvIheeSSTD0yk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        importpath = "go.mongodb.org/mongo-driver",
        sum = "h1:ny3p0reEpgsR2cfA5cjgwFZg3Cv/ofFh/8jbhGtz9VI=",
        version = "v1.7.5",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        sum = "h1:ZvwS0R+56ePWxUNi+Atn9dWONBPp/AUETXlHW0DxSjE=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_uber_go_goleak",
        importpath = "go.uber.org/goleak",
        sum = "h1:2K3zAYmnTNqV73imy9J1T3WC+gmCePx2hEGkimedGto=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:7fIwc/ZtS0q++VgcfqFDxSBZVv/Xo49/SYnDFupUwlI=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_ratelimit",
        importpath = "go.uber.org/ratelimit",
        sum = "h1:K4qVE+byfv/B3tC+4nYWP7v/6SimcO7HzHekoMNBma0=",
        version = "v0.3.1",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        sum = "h1:FiJd5l1UOLj0wCgbSE0rwwXHzEdAZS6hiiSnxJN/D60=",
        version = "v1.24.0",
    )
    go_repository(
        name = "tech_einride_go_aip",
        importpath = "go.einride.tech/aip",
        sum = "h1:bPo4oqBo2ZQeBKo4ZzLb1kxYXTY1ysJhpvQyfuGzvps=",
        version = "v0.73.0",
    )
