// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var locations = []string{
	"Africa/",
	"Africa/Johannesburg",
	"Africa/Kigali",
	"Africa/Algiers",
	"Africa/Tunis",
	"Africa/Casablanca",
	"Africa/Ceuta",
	"Africa/El_Aaiun",
	"Africa/Brazzaville",
	"Africa/Djibouti",
	"Africa/Accra",
	"Africa/Lagos",
	"Africa/Ndjamena",
	"Africa/Tripoli",
	"Africa/Nairobi",
	"Africa/Douala",
	"Africa/Lubumbashi",
	"Africa/Bangui",
	"Africa/Lome",
	"Africa/Timbuktu",
	"Africa/Lusaka",
	"Africa/Addis_Ababa",
	"Africa/Bamako",
	"Africa/Monrovia",
	"Africa/Asmera",
	"Africa/Conakry",
	"Africa/Bissau",
	"Africa/Mbabane",
	"Africa/Bujumbura",
	"Africa/Luanda",
	"Africa/Porto-Novo",
	"Africa/Juba",
	"Africa/Malabo",
	"Africa/Kinshasa",
	"Africa/Dakar",
	"Africa/Nouakchott",
	"Africa/Khartoum",
	"Africa/Harare",
	"Africa/Mogadishu",
	"Africa/Blantyre",
	"Africa/Windhoek",
	"Africa/Libreville",
	"Africa/Freetown",
	"Africa/Sao_Tome",
	"Africa/Asmara",
	"Africa/Abidjan",
	"Africa/Ouagadougou",
	"Africa/Maputo",
	"Africa/Gaborone",
	"Africa/Banjul",
	"Africa/Kampala",
	"Africa/Dar_es_Salaam",
	"Africa/Cairo",
	"Africa/Maseru",
	"Africa/Niamey",
	"America/",
	"America/Halifax",
	"America/Havana",
	"America/Santiago",
	"America/Tegucigalpa",
	"America/Merida",
	"America/Kentucky/",
	"America/Kentucky/Louisville",
	"America/Kentucky/Monticello",
	"America/Louisville",
	"America/Mendoza",
	"America/Catamarca",
	"America/Nipigon",
	"America/Grenada",
	"America/Fortaleza",
	"America/Atikokan",
	"America/Dawson",
	"America/Yellowknife",
	"America/Shiprock",
	"America/Campo_Grande",
	"America/Nassau",
	"America/St_Lucia",
	"America/Antigua",
	"America/Guatemala",
	"America/Ojinaga",
	"America/Juneau",
	"America/Dominica",
	"America/Coral_Harbour",
	"America/Port_of_Spain",
	"America/Belize",
	"America/Edmonton",
	"America/Sao_Paulo",
	"America/St_Barthelemy",
	"America/Marigot",
	"America/Santa_Isabel",
	"America/Guyana",
	"America/El_Salvador",
	"America/Fort_Wayne",
	"America/Belem",
	"America/Barbados",
	"America/Cayenne",
	"America/Indiana/",
	"America/Indiana/Vevay",
	"America/Indiana/Petersburg",
	"America/Indiana/Vincennes",
	"America/Indiana/Marengo",
	"America/Indiana/Tell_City",
	"America/Indiana/Winamac",
	"America/Indiana/Indianapolis",
	"America/Indiana/Knox",
	"America/Miquelon",
	"America/Pangnirtung",
	"America/Anchorage",
	"America/Montevideo",
	"America/Rosario",
	"America/New_York",
	"America/Metlakatla",
	"America/Lima",
	"America/Santo_Domingo",
	"America/Whitehorse",
	"America/Buenos_Aires",
	"America/Creston",
	"America/Boa_Vista",
	"America/Matamoros",
	"America/Vancouver",
	"America/St_Thomas",
	"America/Bahia",
	"America/Monterrey",
	"America/Los_Angeles",
	"America/Detroit",
	"America/Araguaina",
	"America/Glace_Bay",
	"America/Jujuy",
	"America/Rankin_Inlet",
	"America/Thule",
	"America/Port-au-Prince",
	"America/Lower_Princes",
	"America/Mexico_City",
	"America/Mazatlan",
	"America/Atka",
	"America/Toronto",
	"America/Menominee",
	"America/St_Vincent",
	"America/Inuvik",
	"America/Managua",
	"America/Dawson_Creek",
	"America/Noronha",
	"America/Moncton",
	"America/Guayaquil",
	"America/Martinique",
	"America/Guadeloupe",
	"America/Santarem",
	"America/St_Kitts",
	"America/Yakutat",
	"America/Costa_Rica",
	"America/Danmarkshavn",
	"America/Chicago",
	"America/Winnipeg",
	"America/Rainy_River",
	"America/Bahia_Banderas",
	"America/Panama",
	"America/Porto_Velho",
	"America/Paramaribo",
	"America/Ensenada",
	"America/Rio_Branco",
	"America/Asuncion",
	"America/Cordoba",
	"America/Denver",
	"America/Resolute",
	"America/Manaus",
	"America/Swift_Current",
	"America/Recife",
	"America/Hermosillo",
	"America/Scoresbysund",
	"America/Indianapolis",
	"America/Nome",
	"America/Adak",
	"America/Porto_Acre",
	"America/Tortola",
	"America/Cayman",
	"America/Sitka",
	"America/Argentina/",
	"America/Argentina/Mendoza",
	"America/Argentina/Catamarca",
	"America/Argentina/ComodRivadavia",
	"America/Argentina/Ushuaia",
	"America/Argentina/La_Rioja",
	"America/Argentina/Rio_Gallegos",
	"America/Argentina/Buenos_Aires",
	"America/Argentina/San_Luis",
	"America/Argentina/Jujuy",
	"America/Argentina/Tucuman",
	"America/Argentina/San_Juan",
	"America/Argentina/Cordoba",
	"America/Argentina/Salta",
	"America/North_Dakota/",
	"America/North_Dakota/Center",
	"America/North_Dakota/New_Salem",
	"America/North_Dakota/Beulah",
	"America/Iqaluit",
	"America/Cuiaba",
	"America/Montreal",
	"America/Blanc-Sablon",
	"America/Knox_IN",
	"America/Cambridge_Bay",
	"America/Boise",
	"America/St_Johns",
	"America/Puerto_Rico",
	"America/Aruba",
	"America/Curacao",
	"America/Eirunepe",
	"America/Virgin",
	"America/Phoenix",
	"America/Kralendijk",
	"America/Montserrat",
	"America/Godthab",
	"America/Bogota",
	"America/Cancun",
	"America/Anguilla",
	"America/La_Paz",
	"America/Maceio",
	"America/Grand_Turk",
	"America/Caracas",
	"America/Jamaica",
	"America/Regina",
	"America/Goose_Bay",
	"America/Thunder_Bay",
	"America/Tijuana",
	"America/Chihuahua",
	"Antarctica/",
	"Antarctica/Davis",
	"Antarctica/Rothera",
	"Antarctica/Vostok",
	"Antarctica/Syowa",
	"Antarctica/South_Pole",
	"Antarctica/Mawson",
	"Antarctica/DumontDUrville",
	"Antarctica/Troll",
	"Antarctica/McMurdo",
	"Antarctica/Macquarie",
	"Antarctica/Palmer",
	"Antarctica/Casey",
	"Arctic/",
	"Arctic/Longyearbyen",
	"Asia/",
	"Asia/Kolkata",
	"Asia/Muscat",
	"Asia/Hebron",
	"Asia/Yekaterinburg",
	"Asia/Kuwait",
	"Asia/Taipei",
	"Asia/Vladivostok",
	"Asia/Choibalsan",
	"Asia/Aden",
	"Asia/Dhaka",
	"Asia/Ashgabat",
	"Asia/Dili",
	"Asia/Irkutsk",
	"Asia/Baghdad",
	"Asia/Ujung_Pandang",
	"Asia/Macau",
	"Asia/Makassar",
	"Asia/Kabul",
	"Asia/Bahrain",
	"Asia/Pyongyang",
	"Asia/Damascus",
	"Asia/Qyzylorda",
	"Asia/Ust-Nera",
	"Asia/Chongqing",
	"Asia/Tokyo",
	"Asia/Srednekolymsk",
	"Asia/Kathmandu",
	"Asia/Chita",
	"Asia/Manila",
	"Asia/Vientiane",
	"Asia/Ulan_Bator",
	"Asia/Beirut",
	"Asia/Sakhalin",
	"Asia/Omsk",
	"Asia/Yerevan",
	"Asia/Hovd",
	"Asia/Gaza",
	"Asia/Tashkent",
	"Asia/Istanbul",
	"Asia/Almaty",
	"Asia/Oral",
	"Asia/Saigon",
	"Asia/Aqtau",
	"Asia/Jerusalem",
	"Asia/Tel_Aviv",
	"Asia/Yakutsk",
	"Asia/Ashkhabad",
	"Asia/Magadan",
	"Asia/Thimbu",
	"Asia/Riyadh",
	"Asia/Amman",
	"Asia/Dacca",
	"Asia/Urumqi",
	"Asia/Ho_Chi_Minh",
	"Asia/Kamchatka",
	"Asia/Katmandu",
	"Asia/Nicosia",
	"Asia/Krasnoyarsk",
	"Asia/Kashgar",
	"Asia/Harbin",
	"Asia/Kuching",
	"Asia/Singapore",
	"Asia/Thimphu",
	"Asia/Hong_Kong",
	"Asia/Dushanbe",
	"Asia/Samarkand",
	"Asia/Qatar",
	"Asia/Aqtobe",
	"Asia/Bishkek",
	"Asia/Jakarta",
	"Asia/Tehran",
	"Asia/Kuala_Lumpur",
	"Asia/Khandyga",
	"Asia/Ulaanbaatar",
	"Asia/Rangoon",
	"Asia/Brunei",
	"Asia/Anadyr",
	"Asia/Karachi",
	"Asia/Novosibirsk",
	"Asia/Macao",
	"Asia/Jayapura",
	"Asia/Dubai",
	"Asia/Shanghai",
	"Asia/Novokuznetsk",
	"Asia/Bangkok",
	"Asia/Seoul",
	"Asia/Baku",
	"Asia/Tbilisi",
	"Asia/Phnom_Penh",
	"Asia/Colombo",
	"Asia/Pontianak",
	"Asia/Calcutta",
	"Asia/Chungking",
	"Atlantic/",
	"Atlantic/Bermuda",
	"Atlantic/Stanley",
	"Atlantic/Faeroe",
	"Atlantic/Azores",
	"Atlantic/St_Helena",
	"Atlantic/Jan_Mayen",
	"Atlantic/Canary",
	"Atlantic/South_Georgia",
	"Atlantic/Madeira",
	"Atlantic/Reykjavik",
	"Atlantic/Faroe",
	"Atlantic/Cape_Verde",
	"Australia/",
	"Australia/Melbourne",
	"Australia/West",
	"Australia/Perth",
	"Australia/South",
	"Australia/ACT",
	"Australia/Canberra",
	"Australia/Adelaide",
	"Australia/Hobart",
	"Australia/Darwin",
	"Australia/Victoria",
	"Australia/Yancowinna",
	"Australia/Brisbane",
	"Australia/Sydney",
	"Australia/Lindeman",
	"Australia/Tasmania",
	"Australia/Currie",
	"Australia/Queensland",
	"Australia/Lord_Howe",
	"Australia/NSW",
	"Australia/Eucla",
	"Australia/North",
	"Australia/LHI",
	"Australia/Broken_Hill",
	"Brazil/",
	"Brazil/West",
	"Brazil/Acre",
	"Brazil/East",
	"Brazil/DeNoronha",
	"Canada/",
	"Canada/Saskatchewan",
	"Canada/Eastern",
	"Canada/Pacific",
	"Canada/Newfoundland",
	"Canada/Central",
	"Canada/Atlantic",
	"Canada/East-Saskatchewan",
	"Canada/Mountain",
	"Canada/Yukon",
	"CET",
	"Chile/",
	"Chile/EasterIsland",
	"Chile/Continental",
	"CST6CDT",
	"Cuba",
	"EET",
	"Egypt",
	"Eire",
	"EST",
	"EST5EDT",
	"Etc/",
	"Etc/GMT0",
	"Etc/GMT-10",
	"Etc/GMT+8",
	"Etc/Zulu",
	"Etc/GMT-5",
	"Etc/GMT+5",
	"Etc/GMT-12",
	"Etc/GMT+6",
	"Etc/GMT-13",
	"Etc/GMT+10",
	"Etc/GMT-8",
	"Etc/GMT+0",
	"Etc/UCT",
	"Etc/GMT+3",
	"Etc/Universal",
	"Etc/GMT-0",
	"Etc/GMT-1",
	"Etc/GMT+9",
	"Etc/GMT+4",
	"Etc/GMT-14",
	"Etc/GMT-4",
	"Etc/GMT-3",
	"Etc/GMT-9",
	"Etc/GMT+11",
	"Etc/GMT-2",
	"Etc/GMT+12",
	"Etc/GMT+2",
	"Etc/GMT",
	"Etc/Greenwich",
	"Etc/GMT-7",
	"Etc/GMT-11",
	"Etc/GMT-6",
	"Etc/GMT+1",
	"Etc/UTC",
	"Etc/GMT+7",
	"Europe/",
	"Europe/Mariehamn",
	"Europe/Oslo",
	"Europe/Malta",
	"Europe/Berlin",
	"Europe/Dublin",
	"Europe/Riga",
	"Europe/London",
	"Europe/Tiraspol",
	"Europe/Uzhgorod",
	"Europe/Ljubljana",
	"Europe/Vatican",
	"Europe/Jersey",
	"Europe/Warsaw",
	"Europe/Helsinki",
	"Europe/Tallinn",
	"Europe/Andorra",
	"Europe/Zaporozhye",
	"Europe/Tirane",
	"Europe/Minsk",
	"Europe/Vienna",
	"Europe/Amsterdam",
	"Europe/Luxembourg",
	"Europe/Bratislava",
	"Europe/Sofia",
	"Europe/Istanbul",
	"Europe/Monaco",
	"Europe/Chisinau",
	"Europe/Sarajevo",
	"Europe/Moscow",
	"Europe/Bucharest",
	"Europe/Guernsey",
	"Europe/Vaduz",
	"Europe/Simferopol",
	"Europe/Athens",
	"Europe/Nicosia",
	"Europe/Zurich",
	"Europe/Samara",
	"Europe/San_Marino",
	"Europe/Podgorica",
	"Europe/Volgograd",
	"Europe/Brussels",
	"Europe/Gibraltar",
	"Europe/Lisbon",
	"Europe/Kaliningrad",
	"Europe/Copenhagen",
	"Europe/Paris",
	"Europe/Prague",
	"Europe/Madrid",
	"Europe/Skopje",
	"Europe/Belfast",
	"Europe/Belgrade",
	"Europe/Zagreb",
	"Europe/Vilnius",
	"Europe/Isle_of_Man",
	"Europe/Stockholm",
	"Europe/Busingen",
	"Europe/Budapest",
	"Europe/Rome",
	"Europe/Kiev",
	"Factory",
	"GB",
	"GB-Eire",
	"GMT",
	"GMT0",
	"GMT-0",
	"GMT+0",
	"Greenwich",
	"Hongkong",
	"HST",
	"Iceland",
	"Indian/",
	"Indian/Christmas",
	"Indian/Kerguelen",
	"Indian/Mahe",
	"Indian/Chagos",
	"Indian/Maldives",
	"Indian/Reunion",
	"Indian/Comoro",
	"Indian/Mayotte",
	"Indian/Antananarivo",
	"Indian/Mauritius",
	"Indian/Cocos",
	"Iran",
	"Israel",
	"Jamaica",
	"Japan",
	"Kwajalein",
	"Libya",
	"MET",
	"Mexico/",
	"Mexico/BajaNorte",
	"Mexico/BajaSur",
	"Mexico/General",
	"MST",
	"MST7MDT",
	"Navajo",
	"NZ",
	"NZ-CHAT",
	"Pacific/",
	"Pacific/Fakaofo",
	"Pacific/Noumea",
	"Pacific/Efate",
	"Pacific/Galapagos",
	"Pacific/Johnston",
	"Pacific/Pitcairn",
	"Pacific/Marquesas",
	"Pacific/Guam",
	"Pacific/Apia",
	"Pacific/Guadalcanal",
	"Pacific/Kosrae",
	"Pacific/Bougainville",
	"Pacific/Funafuti",
	"Pacific/Niue",
	"Pacific/Kiritimati",
	"Pacific/Nauru",
	"Pacific/Chuuk",
	"Pacific/Enderbury",
	"Pacific/Honolulu",
	"Pacific/Pohnpei",
	"Pacific/Auckland",
	"Pacific/Wake",
	"Pacific/Gambier",
	"Pacific/Easter",
	"Pacific/Midway",
	"Pacific/Majuro",
	"Pacific/Wallis",
	"Pacific/Port_Moresby",
	"Pacific/Tongatapu",
	"Pacific/Saipan",
	"Pacific/Tahiti",
	"Pacific/Norfolk",
	"Pacific/Yap",
	"Pacific/Palau",
	"Pacific/Tarawa",
	"Pacific/Fiji",
	"Pacific/Pago_Pago",
	"Pacific/Samoa",
	"Pacific/Rarotonga",
	"Pacific/Truk",
	"Pacific/Ponape",
	"Pacific/Chatham",
	"Pacific/Kwajalein",
	"Poland",
	"Portugal",
	"PRC",
	"PST8PDT",
	"ROC",
	"ROK",
	"Singapore",
	"Turkey",
	"UCT",
	"Universal",
	"US/",
	"US/Alaska",
	"US/Aleutian",
	"US/Eastern",
	"US/Indiana-Starke",
	"US/Pacific",
	"US/Pacific-New",
	"US/Hawaii",
	"US/Central",
	"US/East-Indiana",
	"US/Arizona",
	"US/Michigan",
	"US/Mountain",
	"US/Samoa",
	"UTC",
	"WET",
	"W-SU",
	"Zulu",
}

// locations_allCmd represents the locations_all command
var locationsAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Prints all location names list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		for i := range locations {
			fmt.Println(locations[i])
		}
	},
}

func init() {
	locationsCmd.AddCommand(locationsAllCmd)
}
