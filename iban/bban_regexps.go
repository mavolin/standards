package iban

import (
	"regexp"

	"github.com/mavolin/standards/iso3166"
)

// Code generated by tools/codegen/bban_regexp; DO NOT EDIT.

var bbanRegexps = map[iso3166.Alpha2Code]*regexp.Regexp{
	iso3166.VA: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.ST: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.BH: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.DE: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.AE: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.XK: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.KW: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.LU: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.NL: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.MR: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.LV: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.GI: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.GL: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d)"),
	iso3166.SK: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<accountNumberPrefix>\\d\\d\\d\\d)(?P<branchCode>\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.CR: regexp.MustCompile("(?:\\d)(?P<bankCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.LY: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<branchCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.CH: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.VG: regexp.MustCompile("(?P<bankCode>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.AL: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<checksum>\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.SC: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z]\\d\\d)(?P<branchCode>\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<currencyCode>[A-Z][A-Z][A-Z])"),
	iso3166.CZ: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<accountNumberPrefix>\\d\\d\\d\\d)(?P<branchCode>\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.JO: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<branchCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.RO: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.QA: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.LT: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.NO: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d)"),
	iso3166.RU: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.MD: regexp.MustCompile("(?P<bankCode>[\\dA-Z][\\dA-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.MT: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.PS: regexp.MustCompile("(?P<bankCode>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.TN: regexp.MustCompile("(?P<bankCode>\\d\\d)(?P<branchCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.HR: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.TL: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.BG: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<branchCode>\\d\\d\\d\\d)(?P<accountType>\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.KZ: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.BR: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<accountType>[A-Z])(?P<ownerAccountNumber>[\\dA-Z])"),
	iso3166.AT: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.DO: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.GR: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.GT: regexp.MustCompile("(?P<bankCode>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<currencyCode>[\\dA-Z][\\dA-Z])(?P<accountType>[\\dA-Z][\\dA-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.EE: regexp.MustCompile("(?P<bankCode>\\d\\d)(?P<branchCode>\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d)"),
	iso3166.GE: regexp.MustCompile("(?P<bankCode>[\\dA-Z][\\dA-Z])(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.LI: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.SE: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d)"),
	iso3166.SD: regexp.MustCompile("(?P<bankCode>\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.IT: regexp.MustCompile("(?P<checksum>[A-Z])(?P<bankCode>\\d\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.LB: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.MK: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<checksum>\\d\\d)"),
	iso3166.LC: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.EG: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.BY: regexp.MustCompile("(?P<bankCode>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<balanceAccountNumber>\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.AD: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.IE: regexp.MustCompile("(?P<bicBankCode>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<bankCode>\\d\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.ME: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.TR: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d)(?:[\\dA-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.IQ: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<branchCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.MU: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z]\\d\\d)(?P<branchCode>\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?:\\d\\d\\d)(?P<currencyCode>[A-Z][A-Z][A-Z])"),
	iso3166.DK: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d)"),
	iso3166.SA: regexp.MustCompile("(?P<bankCode>\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.ES: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<checksum>\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.FR: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<checksum>\\d\\d)"),
	iso3166.RS: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.BA: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<branchCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.MC: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<checksum>\\d\\d)"),
	iso3166.PK: regexp.MustCompile("(?P<bankCode>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.PT: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.BE: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.HU: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<checksum>\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d)"),
	iso3166.GB: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<branchCode>\\d\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.UA: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.IS: regexp.MustCompile("(?P<bankCode>\\d\\d)(?P<branchCode>\\d\\d)(?P<accountType>\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d)(?P<ownerIdentificationNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.SM: regexp.MustCompile("(?P<checksum>[A-Z])(?P<bankCode>\\d\\d\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.FO: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d)"),
	iso3166.SV: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.CY: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d\\d)(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.SI: regexp.MustCompile("(?P<bankCode>\\d\\d)(?P<branchCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d\\d)"),
	iso3166.AZ: regexp.MustCompile("(?P<bankCode>[A-Z][A-Z][A-Z][A-Z])(?P<accountNumber>[\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z][\\dA-Z])"),
	iso3166.IL: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<branchCode>\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.PL: regexp.MustCompile("(?P<bankCode>\\d\\d\\d)(?P<branchCode>\\d\\d\\d\\d)(?P<checksum>\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d\\d)"),
	iso3166.FI: regexp.MustCompile("(?P<bankCode>\\d\\d\\d\\d\\d\\d)(?P<accountNumber>\\d\\d\\d\\d\\d\\d\\d)(?P<checksum>\\d)"),
}