/*
 * MIT License
 *
 * Copyright (C) 2022 Crimson Technologies, LLC. All rights reserved.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package adyen

import "regexp"

var (
	mastercardPattern         = regexp.MustCompile(`^(5[1-5]\d{0,14}|2[2-7]\d{0,14})$`)
	visadankortPattern        = regexp.MustCompile(`^(4571)\d{0,12}$`)
	visaPattern               = regexp.MustCompile(`^4\d{0,18}$`)
	amexPattern               = regexp.MustCompile(`^3[47]\d{0,13}$`)
	dinersPattern             = regexp.MustCompile(`^(36)\d{0,12}$`)
	maestroukPattern          = regexp.MustCompile(`^(6759)\d{0,15}$`)
	soloPattern               = regexp.MustCompile(`^(6767)\d{0,15}$`)
	laserPattern              = regexp.MustCompile(`^(6304|6706|6709|6771)\d{0,15}$`)
	discoverPattern           = regexp.MustCompile(`^(6011\d{0,12}|(644|645|646|647|648|649)\d{0,13}|65\d{0,14})$`)
	jcbPattern                = regexp.MustCompile(`^(352[8,9]\d{0,15}|35[4-8]\d{0,16})$`)
	bcmcPattern               = regexp.MustCompile(`^((6703)\d{0,15}|(479658|606005)\d{0,13})$`)
	bijcardPattern            = regexp.MustCompile(`^(5100081)\d{0,9}$`)
	dankortPattern            = regexp.MustCompile(`^(5019)\d{0,12}$`)
	hipercardPattern          = regexp.MustCompile(`^(606282)\d{0,10}$`)
	cupPattern                = regexp.MustCompile(`^(62|81)\d{0,17}$`)
	maestroPattern            = regexp.MustCompile(`^(5[0|6-8]\d{0,17}|6\d{0,18})$`)
	eloPattern                = regexp.MustCompile(`^((((506699)|(506770)|(506771)|(506772)|(506773)|(506774)|(506775)|(506776)|(506777)|(506778)|(401178)|(438935)|(451416)|(457631)|(457632)|(504175)|(627780)|(636368)|(636297))\d{0,10})|((50676)|(50675)|(50674)|(50673)|(50672)|(50671)|(50670))\d{0,11})$`)
	uatpPattern               = regexp.MustCompile(`^1\d{0,14}$`)
	cartebancairePattern      = regexp.MustCompile(`^[4-6]\d{0,15}$`)
	visaAlphaBankBonusPattern = regexp.MustCompile(`^(450903)\d{0,10}$`)
	mcAlphaBankBonusPattern   = regexp.MustCompile(`^(510099)\d{0,10}$`)
	hiperPattern              = regexp.MustCompile(`^(637095|637568|637599|637609|637612)\d{0,10}$`)
	oasisPattern              = regexp.MustCompile(`^(982616)\d{0,10}$`)
	karenMillenPattern        = regexp.MustCompile(`^(98261465)\d{0,8}$`)
	warehousePattern          = regexp.MustCompile(`^(982633)\d{0,10}$`)
	mirPattern                = regexp.MustCompile(`^(220)\d{0,16}$`)
	codensaPattern            = regexp.MustCompile(`^(590712)\d{0,10}$`)
	naranjaPattern            = regexp.MustCompile(`^(37|40|5[28])([279])\d*$`)
	cabalPattern              = regexp.MustCompile(`^(58|6[03])([03469])\d*$`)
	shoppingPattern           = regexp.MustCompile(`^(27|58|60)([39])\d*$`)
	argenCardPattern          = regexp.MustCompile(`^(50)(1)\d*$`)
	troyPattern               = regexp.MustCompile(`^(97)(9)\d*$`)
	forbrugsforeningenPattern = regexp.MustCompile(`^(60)(0)\d*$`)
	vpayPattern               = regexp.MustCompile(`^(40[1,8]|413|43[4,5]|44[1,23467]|45[5,8]|46[0,136]|47[1,9]|48[2,37])\d{0,16}$`)
	rupayPattern              = regexp.MustCompile(`^(100003|508(2|[5-9])|60(69|[7-8])|652(1[5-9]|[2-5]\d|8[5-9])|65300[3-4]|8172([0-1]|[3-5]|7|9)|817(3[3-8]|40[6-9]|410)|35380([0-2]|[5-6]|9))\d{0,12}$`)
)

// DetectCardType detects the type of the given card number.
// The card number must have no whitespace characters.
//
// If the card's type cannot be detected, then "noBrand" is returned
// which is also what Adyen uses if it cannot detect the card type.
func DetectCardType(formattedCardNumber string) string {
	switch {
	case mastercardPattern.MatchString(formattedCardNumber):
		return "mc"
	case visadankortPattern.MatchString(formattedCardNumber):
		return "visadankort"
	case visaPattern.MatchString(formattedCardNumber):
		return "visa"
	case amexPattern.MatchString(formattedCardNumber):
		return "amex"
	case dinersPattern.MatchString(formattedCardNumber):
		return "diners"
	case maestroukPattern.MatchString(formattedCardNumber):
		return "maestrouk"
	case soloPattern.MatchString(formattedCardNumber):
		return "solo"
	case laserPattern.MatchString(formattedCardNumber):
		return "laser"
	case discoverPattern.MatchString(formattedCardNumber):
		return "discover"
	case jcbPattern.MatchString(formattedCardNumber):
		return "jcb"
	case bcmcPattern.MatchString(formattedCardNumber):
		return "bcmc"
	case bijcardPattern.MatchString(formattedCardNumber):
		return "bijcard"
	case dankortPattern.MatchString(formattedCardNumber):
		return "dankort"
	case hipercardPattern.MatchString(formattedCardNumber):
		return "hiper"
	case cupPattern.MatchString(formattedCardNumber):
		return "cup"
	case maestroPattern.MatchString(formattedCardNumber):
		return "maestro"
	case eloPattern.MatchString(formattedCardNumber):
		return "elo"
	case uatpPattern.MatchString(formattedCardNumber):
		return "uatp"
	case cartebancairePattern.MatchString(formattedCardNumber):
		return "cartebancaire"
	case visaAlphaBankBonusPattern.MatchString(formattedCardNumber):
		return "visaalphabankbonus"
	case mcAlphaBankBonusPattern.MatchString(formattedCardNumber):
		return "mcalphabankbonus"
	case hiperPattern.MatchString(formattedCardNumber):
		return "hiper"
	case oasisPattern.MatchString(formattedCardNumber):
		return "oasis"
	case karenMillenPattern.MatchString(formattedCardNumber):
		return "karenmillen"
	case warehousePattern.MatchString(formattedCardNumber):
		return "warehouse"
	case mirPattern.MatchString(formattedCardNumber):
		return "mir"
	case codensaPattern.MatchString(formattedCardNumber):
		return "codensa"
	case naranjaPattern.MatchString(formattedCardNumber):
		return "naranja"
	case cabalPattern.MatchString(formattedCardNumber):
		return "cabal"
	case shoppingPattern.MatchString(formattedCardNumber):
		return "shopping"
	case argenCardPattern.MatchString(formattedCardNumber):
		return "argencard"
	case troyPattern.MatchString(formattedCardNumber):
		return "troy"
	case forbrugsforeningenPattern.MatchString(formattedCardNumber):
		return "forbrugsforeningen"
	case vpayPattern.MatchString(formattedCardNumber):
		return "vpay"
	case rupayPattern.MatchString(formattedCardNumber):
		return "rupay"
	default:
		return "noBrand"
	}
}
