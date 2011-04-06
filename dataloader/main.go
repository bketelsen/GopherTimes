package main

import (
	"launchpad.net/mgo"
		"log"

)

type Page struct {
	Path    string
	Title   string
	Content string
	Products		[]Product
	PressReleases	[]PressRelease
}


type Product struct {
	Name            string
	Blurb           string
	FullDescription string
	ImagePath       string
}

type PressRelease struct {
	Date      string
	Title     string
	PathToPdf string
}

func main() {

	mongo, err := mgo.Mongo("localhost")
	if err != nil {
		panic(err)
	}

	c := mongo.DB("public_web").C("page")

	p := &Page{Title: "Clarity Services - Products", Path:"products",Content:`  <div id="content" class="inner-page"> 
	        <input type="hidden" name="scale" value="1" /> 

	        <div id="content" class="inner-page"> 

	  <div class="content-text"> 
	    <p>Clarity has a suite of products and tools to help lenders make solid decisions in times of uncertainty:</p> 

	    <table width="100%"> 
	      <tr><td><a href="#clear-warning">Clear Warning</a></td><td><a href="#clear-tradeline">Clear Tradeline</a></td><td><a href="#clear-inquiry">Clear Inquiry</a></td></tr> 
	      <tr><td><a href="#clear-pc-fraud">Clear PC Fraud</a></td><td><a href="#clear-id-fraud">Clear ID Fraud</a></td><td><a href="#clear-id-attributes">Clear ID Attributes</a></td></tr> 
	      <tr><td><a href="#clear-profile">Clear Profile</a></td><td><a href="#clear-composite">Clear Composite</a></td><td><a href="#clear-bank">Clear Bank</a></td></tr> 
	      <tr><td><a href="#clear-check">Clear Check</a></td><td><a href="#clear-bureau">Clear Bureau</a></td><td></td></tr> 
	    </table> 
	  </div> 
	  <br/> 
	  <br/> 


	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-bank"></a><a href="/products/clear_bank">Clear Bank™</a></p> 
	    <p>Clear Bank™ quickly reports all known bank history for a consumer along with NSF history, bank activity, accounts closed/dormant, loan amounts, direct deposit verification on prepaid debit cards, and ACH capabilities by the bank. Clear Bank™ not only reports – it establishes a real picture of a consumer’s payback probability.</p> 
	  </div> 
	  <br/> 
	  <br/> 

	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-bureau"></a><a href="/products/clear_bureau">Clear Bureau™</a></p> 
	    <p>By combining Experian consumer data with Clarity’s consumer information files, Clear Bureau™ empowers financial service providers with supplemental consumer comparison data. Clear Bureau™ offers insight into consumer past credit history behavior in the traditional credit world.</p
	  </div> 
	  <br/> 
	  <br/> 

	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-check"></a><a href="/products/clear_check">Clear Check™</a></p> 
	    <p>Providing check approval similar to that used by major retailers, Clear Check™ offers Check Approval and Check Guarantee providing self-risk protection and protects against loss. Protect your decisions.</p> 
	  </div> 
	  <br/> 
	  <br/> 
	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-composite"></a><a href="/products/clear_composite">Clear Composite™</a></p> 
	    <p>Clarity's Clear Composite™ combines predictive data and returns easy-to-read scores on credit risk and credit fraud risk. This combination of sources provides a clearer picture of your risk.</p> 
	  </div> 
	  <br/> 
	  <br/> 
	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-id-attributes"></a><a href="/products/clear_idattributes">Clear ID Attributes™</a></p> 
	    <p>By providing detailed information about customer identity, Clear ID Attributes™ equips organizations with the next generation of risk management capabilities and the knowledge necessary to make better fraud, marketing, collections, verification, and compliance decisions. Clear ID Attributes™ are derived variables created from both raw and summarized data provided by individual identity elements or a combination of identity elements.</p> 
	  </div> 
	  <br/> 
	  <br/> 
	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-id-fraud"></a><a href="/products/clear_idfraud">Clear ID Fraud™</a></p> 
	    <p>Used to assist in identifying consumers who may be attempting to commit identity fraud, Clarity’s Clear ID Fraud™ evaluates various consumer data points and provides an easy-to-read ID Fraud score and supporting reason codes. Partnering with ID Analytics™, an established leader in identity intelligence, Clear ID Fraud™ scores are calculated using non-traditional identity sources and does not use traditional revivification techniques.</p> 
	    <br/> 
	    <br/> 
	  </div> 
	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-inquiry"></a><a href="/products/clear_inquiry">Clear Inquiry™</a></p> 
	    <p>Using consumer data as submitted to Clarity in recent credit applications, Clear Inquiry™ provides powerful pre-validation flags to assist you in your decision-making process.</p> 
	  </div> 
	  <br/> 
	  <br/> 

	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-pc-fraud"></a><a href="/products/clear_pcfraud">Clear PC Fraud™</a></p> 
	    <p>Clear PC Fraud™ identifies individuals attempting to commit fraud by the use of multiple identities on credit applications submitted from a single computer. Utilizing patented technology, Clear PC Fraud™ captures the digital ‘fingerprint’ (or PCPrint) of an applicant’s computer and archives the data for future comparison. For an unparalleled product combination to help identify and reject fraudulent applicants and applications, try Clear ID Fraud™ and Clear PC Fraud™.</p> 
	  </div> 
	  <br/> 
	  <br/> 

	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-profile"></a><a href="/products/clear_profile">Clear Profile™</a></p> 
	    <p>Clear Profile™ is a powerful collections solution reporting tool that reports the three most current instances of critical consumer demographic data such as: address, phone numbers, employer, financial institution, and email. All data is date-stamped to show when the information was contributed/received.</p> 
	  </div> 
	  <br/> 
	  <br/> 

	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-tradeline"></a><a href="/products/clear_tradeline">Clear Tradeline™</a></p> 
	    <p>Provides traditional Payday Single Payment Lender information, bill pay history, installment loans, credit cards, prepaid debit cards, rent-to-own, and mortgage loans. Clear Tradeline™ provides an in-depth look at credit limits, past due loans, and open balances on all known tradeline loans.</p> 
	  </div> 
	  <br/> 
	  <br/> 

	  <div class="content-text"> 
	    <p class="content-product-header"><a name="clear-warning"></a><a href="/products/clear_warning">Clear Warning™</a></p> 
	    <p>Clear Warning™ provides a simple bank account risk ranking and descriptive response for the type of account found for a reported consumer.</p?
	  </div> 
	</div> 


	      </div><!-- /content --> 
	`}
	err = c.Insert(p)
	if err != nil {
		log.Println(err)
	}

	defer mongo.Close()
}
