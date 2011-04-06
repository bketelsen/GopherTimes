require 'rubygems'
require 'mongo_mapper'
class Page
  include MongoMapper::Document
  connection Mongo::Connection.new('localhost')
  set_database_name 'public_web'
  set_collection_name 'page'
  
  key :path, String
  key :title, String
  key :content, Integer
  key :products, Array
  key :pressreleases, Array
  
  
end

p = Page.create
p.path = "marketing/metrics"
p.title = "Clarity Services, Inc."
p.content = <<EOSTRING
    <div id="content" class="inner-page"> 
        <input type="hidden" name="scale" value="1" /> 
 
        <div id="stat_holder"> 
  <div id="real_time_stats"> 
    
  
 
<div class="stat-container"> 
  <div class="container-left-4"> 
    <div id="marketing_stats_product_denial_clear_idfraud" class="flip-chart-4"> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_product_denial_clear_idfraud_mask_0" class="mask" > 
          <div id="marketing_stats_product_denial_clear_idfraud_holder_0" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_product_denial_clear_idfraud_holder_0', 0);
        
//]]>
</script> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_product_denial_clear_idfraud_mask_1" class="mask" > 
          <div id="marketing_stats_product_denial_clear_idfraud_holder_1" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_product_denial_clear_idfraud_holder_1', 0);
        
//]]>
</script> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_product_denial_clear_idfraud_mask_2" class="mask" > 
          <div id="marketing_stats_product_denial_clear_idfraud_holder_2" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_product_denial_clear_idfraud_holder_2', 1);
        
//]]>
</script> 
      
 
      
        <div class="percent" > 
          <img alt="Flip_chart_percent" src="/images/flip_chart_percent.png?1301006135" /> 
        </div> 
      
 
    </div> 
    <div class="flip-chart-definition-4"> 
      % denied on Clear ID Fraud of the last 200 Clear ID Frauds run.
    </div> 
  </div> 
 
  <div class="container-right-4"> 
    Identifying real identity fraud potential in real time.
  </div> 
</div> 
  
 
<div class="stat-container"> 
  <div class="container-left-4"> 
    <div id="marketing_stats_product_denial_clear_tradeline" class="flip-chart-4"> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_product_denial_clear_tradeline_mask_0" class="mask" > 
          <div id="marketing_stats_product_denial_clear_tradeline_holder_0" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_product_denial_clear_tradeline_holder_0', 0);
        
//]]>
</script> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_product_denial_clear_tradeline_mask_1" class="mask" > 
          <div id="marketing_stats_product_denial_clear_tradeline_holder_1" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_product_denial_clear_tradeline_holder_1', 0);
        
//]]>
</script> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_product_denial_clear_tradeline_mask_2" class="mask" > 
          <div id="marketing_stats_product_denial_clear_tradeline_holder_2" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_product_denial_clear_tradeline_holder_2', 2);
        
//]]>
</script> 
      
 
      
        <div class="percent" > 
          <img alt="Flip_chart_percent" src="/images/flip_chart_percent.png?1301006135" /> 
        </div> 
      
 
    </div> 
    <div class="flip-chart-definition-4"> 
      % denied on Clear Tradeline of the last 200 Tradelines run.
    </div> 
  </div> 
 
  <div class="container-right-4"> 
    We have trade lines reported for many different sub and near prime loan instruments.  
  </div> 
</div> 
  
 
<div class="stat-container"> 
  <div class="container-left-4"> 
    <div id="marketing_stats_highest_micr_ssn" class="flip-chart-4"> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_highest_micr_ssn_mask_0" class="mask" > 
          <div id="marketing_stats_highest_micr_ssn_holder_0" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_highest_micr_ssn_holder_0', 8);
        
//]]>
</script> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_highest_micr_ssn_mask_1" class="mask" > 
          <div id="marketing_stats_highest_micr_ssn_holder_1" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_highest_micr_ssn_holder_1', 6);
        
//]]>
</script> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_highest_micr_ssn_mask_2" class="mask" > 
          <div id="marketing_stats_highest_micr_ssn_holder_2" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_highest_micr_ssn_holder_2', 8);
        
//]]>
</script> 
      
      <!-- IE <= 7 needs the dislpay:inline in order to display properly--> 
        <div id="marketing_stats_highest_micr_ssn_mask_3" class="mask" > 
          <div id="marketing_stats_highest_micr_ssn_holder_3" class="holder"></div> 
        </div> 
 
        <script type="text/javascript"> 
//<![CDATA[
 
          setStripPosition('marketing_stats_highest_micr_ssn_holder_3', 3);
        
//]]>
</script> 
      
 
      
 
    </div> 
    <div class="flip-chart-definition-4"> 
      Highest count of Social Security Numbers associated with 1 bank account in the past 200 inquiries.
    </div> 
  </div> 
 
  <div class="container-right-4"> 
    After you make the loan isn't the best time to find out your customer is sharing that account with 352 of his closest fraudster friends.
  </div> 
</div> 
 
  </div> 
</div> 
 
<div class="scrolling-text-marketing-text"> 
  The most expressive filter language in the industry so that you can craft decision filters on literally anything in our system.
</div> 
 
<div id="marketing_stats_latest_denial_reasons"> 
  <div class="rbroundbox"> 
    <div class="rbtop"><div></div></div> 
    <div class="rbcontent"> 
      <marquee id="marketing_stats_latest_denial_reasons_marquee" scrollAmount="1" scrollDelay="100" direction="up" trueSpeed="true" height="50px"> 
        <b>Deny when primary account is a bad account number.</b><br>clear_bank_reason_codes_has? 'A13'<br><b>Deny when number of SSNs with bank account is 3 or more. </b><br>number_of_ssns_with_bank_account.to_i >= 3<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when primary account may be overdrawn.</b><br>clear_bank_reason_codes_has?('A06')<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when more than 2 SSNs with this bank account.</b><br>number_of_ssns_with_bank_account.to_i > 2<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when high risk bank</b><br>clear_bank_reason_codes_has?('A09')<br><b>Deny when more than 2 SSNs with this bank account.</b><br>number_of_ssns_with_bank_account.to_i > 2<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when tradeline score 700 or less. </b><br>clear_tradeline_score.to_i <= 700<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when more than 1 bank account change in last 30 days.</b><br>bank_idfraud.thirty_days_ago.to_i > 1<br><b>Deny when primary account is not verifiable.</b><br>clear_bank_reason_codes_has?('A03')<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny if last activity by group is in last 30 days.</b><br>inquiry.last_seen_by_group >= 30.days.ago if inquiry.last_seen_by_group<br><b>Deny when primary account is not verifiable.</b><br>clear_bank_reason_codes_has?('A03')<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when primary account is recently opened.</b><br>clear_bank_reason_codes_has?('A14')<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when primary account is not verifiable.</b><br>clear_bank_reason_codes_has?('A03')<br><b>Deny when more than 2 SSNs with this bank account.</b><br>number_of_ssns_with_bank_account.to_i > 2<br><b>Deny when inquiry cluster position is 4 or more. </b><br>current_inquiry_cluster_position.to_i >= 4<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when applicant has been seen by the group in the previous 7 days.</b><br>inquiry.last_seen_by_group >= 7.days.ago if inquiry.last_seen_by_group<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny if last activity by group is in last 30 days.</b><br>inquiry.last_seen_by_group >= 30.days.ago if inquiry.last_seen_by_group<br><b>Deny when primary account is not verifiable.</b><br>clear_bank_reason_codes_has?('A03')<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when more than 2 SSNs with this bank account.</b><br>number_of_ssns_with_bank_account.to_i > 2<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false<br><b>Deny when SSN is not valid.</b><br>social_security_valid == false
      </marquee> 
    </div> <!-- /rbcontent--> 
    <div class="rbbot"><div></div></div> 
  </div><!-- /rbroundbox --> 
</div> 
 
 
 
 
 
      </div><!-- /content --> 
EOSTRING
p.save
