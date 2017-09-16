package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

type fundaObject struct {
	address  string
	price    string
	url      url.URL
	imageURL url.URL
}

type fundaObjects []*fundaObject

type fundaSearchResult struct {
	// AccountStatus     int         `json:"AccountStatus"`
	// EmailNotConfirmed bool        `json:"EmailNotConfirmed"`
	// ValidationFailed  bool        `json:"ValidationFailed"`
	// ValidationReport  interface{} `json:"ValidationReport"`
	// Website           int         `json:"Website"`
	// Metadata struct {
	// ObjectType   string `json:"ObjectType"`
	// Omschrijving string `json:"Omschrijving"`
	// Titel        string `json:"Titel"`
	// } `json:"Metadata"`
	Objects []struct {
		// AangebodenSindsTekst        string        `json:"AangebodenSindsTekst"`
		// AanmeldDatum                string        `json:"AanmeldDatum"`
		// AantalBeschikbaar           interface{}   `json:"AantalBeschikbaar"`
		// AantalKamers int `json:"AantalKamers"`
		// AantalKavels                interface{}   `json:"AantalKavels"`
		// Aanvaarding                 string        `json:"Aanvaarding"`
		Adres string `json:"Adres"`
		// Afstand                     int           `json:"Afstand"`
		// BronCode                    string        `json:"BronCode"`
		// ChildrenObjects             []interface{} `json:"ChildrenObjects"`
		// DatumAanvaarding            interface{}   `json:"DatumAanvaarding"`
		// DatumOndertekeningAkte      interface{}   `json:"DatumOndertekeningAkte"`
		// Foto                        string        `json:"Foto"`
		// FotoLarge                   string        `json:"FotoLarge"`
		// FotoLargest                 string        `json:"FotoLargest"`
		FotoMedium string `json:"FotoMedium"`
		// FotoSecure                  string        `json:"FotoSecure"`
		// GewijzigdDatum              interface{}   `json:"GewijzigdDatum"`
		// GlobalID                    int           `json:"GlobalId"`
		// GroupByObjectType           string        `json:"GroupByObjectType"`
		// Heeft360GradenFoto          bool          `json:"Heeft360GradenFoto"`
		// HeeftBrochure               bool          `json:"HeeftBrochure"`
		// HeeftOpenhuizenTopper       bool          `json:"HeeftOpenhuizenTopper"`
		// HeeftOverbruggingsgrarantie bool          `json:"HeeftOverbruggingsgrarantie"`
		// HeeftPlattegrond            bool          `json:"HeeftPlattegrond"`
		// HeeftTophuis                bool          `json:"HeeftTophuis"`
		// HeeftVeiling                bool          `json:"HeeftVeiling"`
		// HeeftVideo                  bool          `json:"HeeftVideo"`
		// HuurPrijsTot                interface{}   `json:"HuurPrijsTot"`
		// Huurprijs                   interface{}   `json:"Huurprijs"`
		// HuurprijsFormaat            interface{}   `json:"HuurprijsFormaat"`
		// ID                          string        `json:"Id"`
		// InUnitsVanaf                interface{}   `json:"InUnitsVanaf"`
		// IndProjectObjectType        bool          `json:"IndProjectObjectType"`
		// IndTransactieMakelaarTonen  interface{}   `json:"IndTransactieMakelaarTonen"`
		// IsSearchable                bool          `json:"IsSearchable"`
		// IsVerhuurd                  bool          `json:"IsVerhuurd"`
		// IsVerkocht                  bool          `json:"IsVerkocht"`
		// IsVerkochtOfVerhuurd        bool          `json:"IsVerkochtOfVerhuurd"`
		// Koopprijs                   int           `json:"Koopprijs"`
		// KoopprijsFormaat            string        `json:"KoopprijsFormaat"`
		// KoopprijsTot                int           `json:"KoopprijsTot"`
		// MakelaarID                  int           `json:"MakelaarId"`
		// MakelaarNaam                string        `json:"MakelaarNaam"`
		// MobileURL                   string        `json:"MobileURL"`
		// Note                        interface{}   `json:"Note"`
		// OpenHuis                    []string      `json:"OpenHuis"`
		// Oppervlakte int `json:"Oppervlakte"`
		// Perceeloppervlakte          interface{}   `json:"Perceeloppervlakte"`
		// Postcode string `json:"Postcode"`
		Prijs struct {
			// GeenExtraKosten     bool        `json:"GeenExtraKosten"`
			// HuurAbbreviation    string      `json:"HuurAbbreviation"`
			// Huurprijs           interface{} `json:"Huurprijs"`
			// HuurprijsOpAanvraag string      `json:"HuurprijsOpAanvraag"`
			// HuurprijsTot        interface{} `json:"HuurprijsTot"`
			KoopAbbreviation string `json:"KoopAbbreviation"`
			Koopprijs        int    `json:"Koopprijs"`
			// KoopprijsOpAanvraag string      `json:"KoopprijsOpAanvraag"`
			// KoopprijsTot        int         `json:"KoopprijsTot"`
			// OriginelePrijs      interface{} `json:"OriginelePrijs"`
			// VeilingText         string      `json:"VeilingText"`
		} `json:"Prijs"`
		// PrijsGeformatteerdHTML     string   `json:"PrijsGeformatteerdHtml"`
		// PrijsGeformatteerdTextHuur string   `json:"PrijsGeformatteerdTextHuur"`
		// PrijsGeformatteerdTextKoop string   `json:"PrijsGeformatteerdTextKoop"`
		// Producten                  []string `json:"Producten"`
		// Project                    struct {
		// 	AantalKamersTotEnMet interface{}   `json:"AantalKamersTotEnMet"`
		// 	AantalKamersVan      interface{}   `json:"AantalKamersVan"`
		// 	AantalKavels         interface{}   `json:"AantalKavels"`
		// 	Adres                interface{}   `json:"Adres"`
		// 	FriendlyURL          interface{}   `json:"FriendlyUrl"`
		// 	GewijzigdDatum       interface{}   `json:"GewijzigdDatum"`
		// 	GlobalID             interface{}   `json:"GlobalId"`
		// 	HoofdFoto            string        `json:"HoofdFoto"`
		// 	IndIpix              bool          `json:"IndIpix"`
		// 	IndPDF               bool          `json:"IndPDF"`
		// 	IndPlattegrond       bool          `json:"IndPlattegrond"`
		// 	IndTop               bool          `json:"IndTop"`
		// 	IndVideo             bool          `json:"IndVideo"`
		// 	InternalID           string        `json:"InternalId"`
		// 	MaxWoonoppervlakte   interface{}   `json:"MaxWoonoppervlakte"`
		// 	MinWoonoppervlakte   interface{}   `json:"MinWoonoppervlakte"`
		// 	Naam                 interface{}   `json:"Naam"`
		// 	Omschrijving         interface{}   `json:"Omschrijving"`
		// 	OpenHuizen           []interface{} `json:"OpenHuizen"`
		// 	Plaats               interface{}   `json:"Plaats"`
		// 	Prijs                interface{}   `json:"Prijs"`
		// 	PrijsGeformatteerd   interface{}   `json:"PrijsGeformatteerd"`
		// 	PublicatieDatum      interface{}   `json:"PublicatieDatum"`
		// 	Type                 int           `json:"Type"`
		// 	Woningtypen          interface{}   `json:"Woningtypen"`
		// } `json:"Project"`
		// ProjectNaam interface{} `json:"ProjectNaam"`
		// PromoLabel  struct {
		// 	HasPromotionLabel     bool          `json:"HasPromotionLabel"`
		// 	PromotionPhotos       []interface{} `json:"PromotionPhotos"`
		// 	PromotionPhotosSecure interface{}   `json:"PromotionPhotosSecure"`
		// 	PromotionType         int           `json:"PromotionType"`
		// 	RibbonColor           int           `json:"RibbonColor"`
		// 	RibbonText            interface{}   `json:"RibbonText"`
		// 	Tagline               interface{}   `json:"Tagline"`
		// } `json:"PromoLabel"`
		// PublicatieDatum string `json:"PublicatieDatum"`
		// PublicatieStatus       int         `json:"PublicatieStatus"`
		// SavedDate              interface{} `json:"SavedDate"`
		// SoortAanbod            string      `json:"Soort-aanbod"`
		// SoortAanbod            int         `json:"SoortAanbod"`
		// StartOplevering        interface{} `json:"StartOplevering"`
		// TimeAgoText            interface{} `json:"TimeAgoText"`
		// TransactieAfmeldDatum  interface{} `json:"TransactieAfmeldDatum"`
		// TransactieMakelaarID   interface{} `json:"TransactieMakelaarId"`
		// TransactieMakelaarNaam interface{} `json:"TransactieMakelaarNaam"`
		// TypeProject            int         `json:"TypeProject"`
		URL string `json:"URL"`
		// VerkoopStatus string `json:"VerkoopStatus"`
		// WGS84X                 float64     `json:"WGS84_X"`
		// WGS84Y                 float64     `json:"WGS84_Y"`
		// WoonOppervlakteTot     int         `json:"WoonOppervlakteTot"`
		// Woonoppervlakte int    `json:"Woonoppervlakte"`
		// Woonplaats      string `json:"Woonplaats"`
		// ZoekType        []int  `json:"ZoekType"`
	} `json:"Objects"`
	Paging struct {
		AantalPaginas int `json:"AantalPaginas"`
		// HuidigePagina int         `json:"HuidigePagina"`
		// VolgendeURL   interface{} `json:"VolgendeUrl"`
		// VorigeURL     interface{} `json:"VorigeUrl"`
	} `json:"Paging"`
	TotaalAantalObjecten int `json:"TotaalAantalObjecten"`
}

func fundaObjectsFromSearchResult(r io.Reader) (objects fundaObjects, pageCount int, err error) {
	var result fundaSearchResult
	if err = json.NewDecoder(r).Decode(&result); err != nil {
		return
	}

	pageCount = result.Paging.AantalPaginas

	for _, o := range result.Objects {
		var houseURL, imageURL *url.URL

		object := &fundaObject{}
		object.address = o.Adres
		object.price = strings.Trim(fmt.Sprintf(
			"€ %v %v",
			humanize.FormatInteger("#.###,", o.Prijs.Koopprijs),
			o.Prijs.KoopAbbreviation,
		), " ")

		houseURL, err = url.Parse(o.URL)
		if err != nil {
			log.Printf("Error parsing house URL: %s", err)
			return
		}
		object.url = *houseURL

		imageURL, err = url.Parse(o.FotoMedium)
		if err != nil {
			log.Printf("Error parsing image URL: %s", err)
			return
		}
		object.imageURL = *imageURL

		objects = append(objects, object)
	}

	return
}

func fundaSearchURL(token, searchOptions string, page, pageSize int) (*url.URL, error) {
	u, err := url.Parse("http://partnerapi.funda.nl/feeds/Aanbod.svc/search/json/" + token + "/")
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Set("website", "funda")
	q.Set("type", "koop")
	q.Set("zo", searchOptions)
	q.Set("page", strconv.Itoa(page))
	q.Set("pagesize", strconv.Itoa(pageSize))

	u.RawQuery = q.Encode()

	return u, nil
}
