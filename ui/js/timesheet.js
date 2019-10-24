function showSummary(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    date = params.get("date_summary");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;

    var request = new XMLHttpRequest();
    var url = "/showSummaryTimesheet";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) { 
            var json = JSON.parse(request.responseText);
            var siamChamnankit = "";
            var countSiamChamnankit = 0 ;
            var totalCoachingSiamChamnankit = 0;
            var totalTraningSiamChamnankit = 0;
            var totalOtherSiamChamnankit = 0;
            var totalIncomesSiamChamnankit = 0;
            var totalSalarySiamChamnankit = 0;
            var totalIncomeTax1SiamChamnankit = 0;
            var totalSocialSecuritySiamChamnankit = 0;
            var totalNetSalarySiamChamnankit = 0;
            var totalWageSiamChamnankit = 0;
            var totalIncomeTax53SiamChamnankit = 0;
            var totalNetWageSiamChamnankit = 0;
            var totalNetTransferSiamChamnankit = 0;
            
            var shuhari = "";
            var countShuhari = 0;
            var totalCoachingShuhari = 0;
            var totalTraningShuhari = 0;
            var totalOtherShuhari = 0;
            var totalIncomesShuhari = 0;
            var totalSalaryShuhari = 0;
            var totalIncomeTax1Shuhari = 0;
            var totalSocialSecurityShuhari = 0;
            var totalNetSalaryShuhari = 0;
            var totalWageShuhari = 0;
            var totalIncomeTax53Shuhari = 0;
            var totalNetWageShuhari = 0;
            var totalNetTransferShuhari = 0;
   
            for (var i = 1; i <= json.length; i++) {  
                if(json[i-1].company === "siam_chamnankit"){
                    countSiamChamnankit++;
                    siamChamnankit += "<tr id=\"row_summary_id_"+i+"\">";
                    siamChamnankit += "<td>"+countSiamChamnankit+"</td>";
                    siamChamnankit += "<td id=\"member_name_th_id_"+i+"\">"+json[i-1].member_name_th+"</td>";
                    siamChamnankit += "<td id=\"coaching_id_"+i+"\">"+json[i-1].coaching+"</td>";
                    totalCoachingSiamChamnankit += json[i-1].coaching;
                    siamChamnankit += "<td id=\"training_id_"+i+"\">"+json[i-1].training+"</td>";
                    totalTraningSiamChamnankit +=json[i-1].training;
                    siamChamnankit += "<td id=\"other_id_"+i+"\">"+json[i-1].other+"</td>";
                    totalOtherSiamChamnankit += json[i-1].other;
                    siamChamnankit += "<td id=\"total_incomes_id_"+i+"\">"+json[i-1].total_incomes+"</td>";
                    totalIncomesSiamChamnankit += json[i-1].total_incomes;
                    siamChamnankit += "<td id=\"salary_id_"+i+"\">"+json[i-1].salary+"</td>";
                    totalSalarySiamChamnankit += json[i-1].salary;
                    siamChamnankit += "<td id=\"income_tax_1_id_"+i+"\">"+json[i-1].income_tax_1+"</td>";
                    totalIncomeTax1SiamChamnankit += json[i-1].income_tax_1;
                    siamChamnankit += "<td id=\"social_security_id_"+i+"\">"+json[i-1].social_security+"</td>";
                    totalSocialSecuritySiamChamnankit += json[i-1].social_security;
                    siamChamnankit += "<td id=\"net_salary_id_"+i+"\">"+json[i-1].net_salary+"</td>";
                    totalNetSalarySiamChamnankit += json[i-1].net_salary;
                    siamChamnankit += "<td id=\"wage_id_"+i+"\">"+json[i-1].wage+"</td>";
                    totalWageSiamChamnankit += json[i-1].wage;
                    siamChamnankit += "<td id=\"income_tax_53_percentage_id_"+i+"\">"+json[i-1].income_tax_53_percentage+"</td>";
                    siamChamnankit += "<td id=\"income_tax_53_id_"+i+"\">"+json[i-1].income_tax_53+"</td>";
                    totalIncomeTax53SiamChamnankit += json[i-1].income_tax_53;
                    siamChamnankit += "<td id=\"net_wage_id_"+i+"\">"+json[i-1].net_wage+"</td>";
                    totalNetWageSiamChamnankit +=json[i-1].net_wage;
                    siamChamnankit += "<td id=\"net_transfer_id_"+i+"\">"+json[i-1].net_transfer+"</td>";
                    totalNetTransferSiamChamnankit += json[i-1].net_transfer;
                    siamChamnankit += "<td><select id=\"status_checking_transfer_"+i+"\"><option value=\""+json[i-1].status_checking_transfer+"\">"+json[i-1].status_checking_transfer+"</option>";
                    siamChamnankit += "<option value=\"รอการตรวจสอบ\">รอการตรวจสอบ</option>";
                    siamChamnankit += "<option value=\"โอนเงินเรียบร้อย\">โอนเงินเรียบร้อย</option>";
                    siamChamnankit += "<option value=\"ถูกต้อง\">ถูกต้อง</option>";
                    siamChamnankit += "<option value=\"ไม่ถูกต้อง\">ไม่ถูกต้อง</option>";
                    siamChamnankit += "</select></td>";
                    siamChamnankit += "<td><input type=\"text\" id=\"date_transfer_"+i+"\" value=\""+json[i-1].date_transfer+"\"></td>";
                    siamChamnankit += "<td><input type=\"text\" id=\"comment_"+i+"\" value=\""+json[i-1].comment+"\"></td>";
                    siamChamnankit += "<input type=\"hidden\" id=\"transaction_id_"+i+"\" value=\""+json[i-1].id+"\">";
                    siamChamnankit += "<td><input type=\"submit\" id=\"button_change_status_checking_transfer_id_"+i+"\" value=\"เปลี่ยนสถานะ\" onclick=\"updateStatusTransfer("+i+")\"/>"+"</td>";
                    siamChamnankit += "</tr>";
            
                }else{
                    countShuhari++;
                    shuhari += "<tr id=\"row_summary_id_"+i+"\">";
                    shuhari += "<td>"+countShuhari+"</td>";
                    shuhari += "<td id=\"member_name_th_id_"+i+"\">"+json[i-1].member_name_th+"</td>";
                    shuhari += "<td id=\"coaching_id_"+i+"\">"+json[i-1].coaching+"</td>";
                    totalCoachingShuhari += json[i-1].coaching;
                    shuhari += "<td id=\"training_id_"+i+"\">"+json[i-1].training+"</td>";
                    totalTraningShuhari += json[i-1].training;
                    shuhari += "<td id=\"other_id_"+i+"\">"+json[i-1].other+"</td>";
                    totalOtherShuhari += json[i-1].other;
                    shuhari += "<td id=\"total_incomes_id_"+i+"\">"+json[i-1].total_incomes+"</td>";
                    totalIncomesShuhari += json[i-1].total_incomes;
                    shuhari += "<td id=\"salary_id_"+i+"\">"+json[i-1].salary+"</td>";
                    totalSalaryShuhari += json[i-1].salary;
                    shuhari += "<td id=\"income_tax_1_id_"+i+"\">"+json[i-1].income_tax_1+"</td>";
                    totalIncomeTax1Shuhari += json[i-1].income_tax_1;
                    shuhari += "<td id=\"social_security_id_"+i+"\">"+json[i-1].social_security+"</td>";
                    totalSocialSecurityShuhari += json[i-1].social_security;
                    shuhari += "<td id=\"net_salary_id_"+i+"\">"+json[i-1].net_salary+"</td>";
                    totalNetSalaryShuhari += json[i-1].net_salary;
                    shuhari += "<td id=\"wage_id_"+i+"\">"+json[i-1].wage+"</td>";
                    totalWageShuhari += json[i-1].wage;
                    shuhari += "<td id=\"income_tax_53_percentage_id_"+i+"\">"+json[i-1].income_tax_53_percentage+"</td>";
                    shuhari += "<td id=\"income_tax_53_id_"+i+"\">"+json[i-1].income_tax_53+"</td>";
                    totalIncomeTax53Shuhari += json[i-1].income_tax_53;
                    shuhari += "<td id=\"net_wage_id_"+i+"\">"+json[i-1].net_wage+"</td>";
                    totalNetWageShuhari += json[i-1].net_wage;
                    shuhari += "<td id=\"net_transfer_id_"+i+"\">"+json[i-1].net_transfer+"</td>";
                    totalNetTransferShuhari += json[i-1].net_transfer;
                    shuhari += "<td><select id=\"status_checking_transfer_"+i+"\"><option value=\""+json[i-1].status_checking_transfer+"\">"+json[i-1].status_checking_transfer+"</option>";
                    shuhari += "<option value=\"รอการตรวจสอบ\">รอการตรวจสอบ</option>";
                    shuhari += "<option value=\"โอนเงินเรียบร้อย\">โอนเงินเรียบร้อย</option>";
                    shuhari += "<option value=\"ถูกต้อง\">ถูกต้อง</option>";
                    shuhari += "<option value=\"ไม่ถูกต้อง\">ไม่ถูกต้อง</option>";
                    shuhari += "</select></td>";
                    shuhari += "<td><input type=\"text\" id=\"date_transfer_"+i+"\" value=\""+json[i-1].date_transfer+"\"></td>";
                    shuhari += "<td><input type=\"text\" id=\"comment_"+i+"\" value=\""+json[i-1].comment+"\"></td>";
                    shuhari += "<input type=\"hidden\" id=\"transaction_id_"+i+"\" value=\""+json[i-1].id+"\">";
                    shuhari += "<td>"+"<input type=\"submit\" id=\"button_change_status_checking_transfer_id_"+i+"\" value=\"เปลี่ยนสถานะ\" onclick=\"updateStatusTransfer("+i+")\"/>"+"</td>";
                    shuhari += "</tr>";
                }
            }
            $("#table_siam_chamnankit").html(siamChamnankit);
            $("#total_coaching_siam_chamnankit").html(totalCoachingSiamChamnankit);
            $("#total_traning_siam_chamnankit").html(totalTraningSiamChamnankit);
            $("#total_other_siam_chamnankit").html(totalOtherSiamChamnankit);
            $("#total_incomes_siam_chamnankit").html(totalIncomesSiamChamnankit);
            $("#total_salary_siam_chamnankit").html(totalSalarySiamChamnankit);
            $("#total_income_tax_1_siam_chamnankit").html(totalIncomeTax1SiamChamnankit);
            $("#total_social_security_siam_chamnankit").html(totalSocialSecuritySiamChamnankit);
            $("#total_net_salary_siam_chamnankit").html(totalNetSalarySiamChamnankit);
            $("#total_wage_siam_chamnankit").html(totalWageSiamChamnankit);
            $("#total_income_tax_53_siam_chamnankit").html(totalIncomeTax53SiamChamnankit);
            $("#total_net_wage_siam_chamnankit").html(totalNetWageSiamChamnankit);
            $("#total_net_transfer_siam_chamnankit").html(totalNetTransferSiamChamnankit);
            
            $("#table_shuhari").html(shuhari);
            $("#total_coaching_shuhari").html(totalCoachingShuhari);
            $("#total_traning_shuhari").html(totalTraningShuhari);
            $("#total_other_shuhari").html(totalOtherShuhari);
            $("#total_incomes_shuhari").html(totalIncomesShuhari);
            $("#total_salary_shuhari").html(totalSalaryShuhari);
            $("#total_income_tax_1_shuhari").html(totalIncomeTax1Shuhari);
            $("#total_social_security_shuhari").html(totalSocialSecurityShuhari);
            $("#total_net_salary_shuhari").html(totalNetSalaryShuhari);
            $("#total_wage_shuhari").html(totalWageShuhari);
            $("#total_income_tax_53_shuhari").html(totalIncomeTax53Shuhari);
            $("#total_net_wage_shuhari").html(totalNetWageShuhari);
            $("#total_net_transferShuhari").html(totalNetTransferShuhari);  
        }
    }; 
    var data = JSON.stringify({"year":year, "month": month});
    request.send(data);  

}

function updateStatusTransfer(index){
    var transactionID = $("#transaction_id_"+index).val();
    var statusTransfer = $("#status_checking_transfer_"+index).val();
    var dateTransfer = $("#date_transfer_"+index).val();
    var comment = $("#comment_"+index).val();

    var request = new XMLHttpRequest();
    var url = "/updateStatusCheckingTransfer";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
        }
    }
    var data = JSON.stringify({"transaction_id":transactionID,"status":statusTransfer,"date":dateTransfer,"comment":comment});
    request.send(data); 
    window.location.replace(window.location.href)    
}

function calculateTimesheet(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    memberID = params.get("id");
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;
    var day = parseInt($("#day").val());

    var fullStartTimeAm = $("#start_time_am").val();
    var startTimeAm = new Date("January 02, 2006 "+fullStartTimeAm+" UTC");
    
    var fullEndTimeAm = $("#end_time_am").val();
    var endTimeAm = new Date("January 02, 2006 "+fullEndTimeAm+" UTC");

    var fullStartTimePm = $("#start_time_pm").val();
    var startTimePm = new Date("January 02, 2006 "+fullStartTimePm+" UTC");

    var fullEndTimePm = $("#end_time_pm").val();
    var endTimePm = new Date("January 02, 2006 "+fullEndTimePm+" UTC");

    var overtime = parseInt($("#overtime").val());

    var fullTotalHours = $("#total_hours").val();
    var totalHours = new Date("January 02, 2006 "+fullTotalHours+" UTC");

    var coachingCustomerCharging = parseFloat($("#coaching_customer_charging").val());

    var coachingPaymentRate = parseFloat($("#coaching_payment_rate").val());
    
    var trainingWage = parseFloat($("#training_wage").val());

    var otherWage = parseFloat($("#other_wage").val());

    var company = $("#company").val();

    var description = $("#description").val();
    
    var request = new XMLHttpRequest();
    var url = "/addIncomeItem";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    var data = JSON.stringify({"year":year,"month":month,"member_id":memberID,"incomes":{"day":day,"start_time_am":startTimeAm,"end_time_am":endTimeAm,"start_time_pm":startTimePm,"end_time_pm":endTimePm,"overtime":overtime,"total_hours":totalHours,"coaching_customer_charging":coachingCustomerCharging,"coaching_payment_rate":coachingPaymentRate,"training_wage":trainingWage,"other_wage":otherWage,"company":company,"description":description}});
    request.send(data);  
    window.location.replace(window.location.href); 
}

function calculatePayment() {
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    memberID = params.get("id");
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;

    var request = new XMLHttpRequest();
    var url = "/calculatePayment";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
        }
    }
    var data = JSON.stringify({"member_id":memberID,"year":year,"month":month});
    request.send(data); 
    window.location.replace(window.location.href); 
}

function showSummaryByID() {
    const monthNames = ["January", "February", "March", "April", "May", "June","July", "August", "September", "October", "November", "December"];

    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    memberID = params.get("id");
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;

    if (date == null || memberID == null) {
        alert("โปรดกรอกข้อมูลให้ครบถ้วน");
        location.href = document.referrer
    }

    var request = new XMLHttpRequest();
    var url = "/showTimesheetByID";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var memberNameENG = json.member_name_eng;
		    var email = json.email;
		    var overtimeRate = json.overtime_rate;
		    var ratePerDay = json.rate_per_day;
            var ratePerHour = json.rate_per_hour;
            var monthString = monthNames[month-1];
		    var totalHours = json.total_hours;
		    var totalCoachingCustomerCharging = json.total_coaching_customer_charging;
            var totalCoachingPaymentRate = json.total_coaching_payment_rate;
            var totalTrainigWage = json.total_training_wage;
            var totalOtherWage = json.total_other_wage;
            var paymentWage = json.payment_wage;
            var incomeList = "";
            
            if (json.incomes !== null) {
                for (var i = 0; i < json.incomes.length; i++) {
                    incomeList += "<tr id=\"income_company_"+json.incomes[i].company+"\">";
                    incomeList += "<td>"+json.incomes[i].day+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].start_time_am)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].end_time_am)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].start_time_pm)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].end_time_pm)+"</td>";
                    incomeList += "<td>"+json.incomes[i].overtime+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].total_hours)+"</td>";
                    incomeList += "<td>"+json.incomes[i].coaching_customer_charging+"</td>";
                    incomeList += "<td>"+json.incomes[i].coaching_payment_rate+"</td>";
                    incomeList += "<td>"+json.incomes[i].training_wage+"</td>";
                    incomeList += "<td>"+json.incomes[i].other_wage+"</td>";
                    incomeList += "<td>"+json.incomes[i].description+"</td>";
                    incomeList += "<td><input type=\"hidden\" id=\"income_id_"+i+"\" value=\""+json.incomes[i].id+"\">"
                    incomeList += "<input id=\"delete_button\" type=\"submit\" value=\"ลบ\" onclick=\"deleteIncome("+i+")\"/>"+"</td>";                    
                    incomeList += "</tr>";
                }
                $("#table_timesheet").html(incomeList);
            }
            $("#member_name_eng").html(memberNameENG);
            $("#email").html(email);
            $("#overtime_rate").html(overtimeRate);
            $("#rate_per_day").html(ratePerDay);
            $("#rate_per_hour").html(ratePerHour);
            $("#month").html(month);
            $("#full_month").html(monthString);
            $("#year").html(year); 
            $("#thours").html(totalHours);
            $("#total_coaching_customer_charging").html(totalCoachingCustomerCharging);
            $("#total_coaching_payment_rate").html(totalCoachingPaymentRate);
            $("#total_trainig_wage").html(totalTrainigWage); 
            $("#total_other_wage").html(totalOtherWage); 
            $("#payment_wage").html(paymentWage); 
                    
        }
    }
    var data = JSON.stringify({"member_id":memberID,"year":year,"month":month});
    request.send(data); 

}

function convertTimestampToTime(timestamp){
    var date = new Date(timestamp);
    datetext = date.toUTCString();
    datetext = datetext.split(' ')[4];
    return datetext
}

function deleteIncome(index){
    var incomeID = parseInt($("#income_id_"+index).val());    

    var request = new XMLHttpRequest();
    var url = "/deleteIncomeItem";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
        }
    }
    var data = JSON.stringify({"id":incomeID});
    
    request.send(data);
    window.location.replace(window.location.href) 
}

function getMemberByID(){
    var urlString = document.referrer;
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    var memberID =  params.get("id");
    
    var request = new XMLHttpRequest();
    var url = "/showMemberDetailsByID";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var member = "";
            for (var i = 0; i < json.length; i++) {
                member += "<table>"
                member += "<tr><th>บริษัท</th><td>"+json[i].company+"</td></tr>";
                member += "<tr><th>ชื่อ(ภาษาไทย)</th><td><input type=\"text\" id=\"member_name_th_id_"+i+"\" value=\""+json[i].member_name_th+"\"></td></tr>";
                member += "<tr><th>ชื่อ(ภาษาอังกฤษ)</th><td><input type=\"text\" id=\"member_name_eng_id_"+i+"\" value=\""+json[i].member_name_eng+"\"></td></tr>";
                member += "<tr><th>E-mail</th><td><input type=\"email\" id=\"email_id_"+i+"\" value=\""+json[i].email+"\"></td></tr>";
                member += "<tr><th>Overtime Rate</th><td><input type=\"number\" id=\"overtime_rate_id_"+i+"\" value=\""+json[i].overtime_rate+"\"></td></tr>";
                member += "<tr><th>Rate Per Day</th><td><input type=\"number\" id=\"rate_per_day_id_"+i+"\" value=\""+json[i].rate_per_day+"\"></td></tr>";
                member += "<tr><th>Rate Per Hour</th><td><input type=\"number\" id=\"rate_per_hour_id_"+i+"\" value=\""+json[i].rate_per_hour+"\"></td></tr>";
                member += "<tr><th>เงินเดือน</th><td><input type=\"number\" id=\"salary_id_"+i+"\" value=\""+json[i].salary+"\"></td></tr>";
                member += "<tr><th>หัก ณ ที่จ่าย ภ.ง.ด.1</th><td><input type=\"number\" id=\"income_tax_1_id_"+i+"\" value=\""+json[i].income_tax_1+"\"></td></tr>";
                member += "<tr><th>ประกันสังคม</th><td><input type=\"number\" id=\"social_security_id_"+i+"\" value=\""+json[i].social_security+"\"></td></tr>";
                member += "<tr><th>หัก ณ ที่จ่าย ภ.ง.ด.53 (ร้อยละ)</th><td><input type=\"number\" id=\"income_tax_53_percentage_id_"+i+"\" value=\""+json[i].income_tax_53_percentage+"\"></td></tr>";
                member += "<tr><th>ประเภทของรายได้</th><td><select id=\"status_id_"+i+"\">";
                member += "<option value=\""+json[i].status+"\">"+json[i].status+"</option>";
                member += "<option value=\"wage\">ค่าจ้างรายวัน (wage)</option>";
                member += "<option value=\"salary\">เงินเดือน (salary)</option>";
                member += "</select></td></tr>";
                member += "<tr><th>ค่าเดินทาง</th><td><input type=\"number\" id=\"travel_expense_id_"+i+"\" value=\""+json[i].travel_expense+"\"></td></tr>";
                member += "<input type=\"hidden\" id=\"member_details_id_"+i+"\" value=\""+json[i].id+"\">";
                member += "<tr><td colspan=\"2\"><input type=\"submit\" id=\"button_edit_member_id_"+i+"\" value=\"ยืนยันการแก้ไขข้อมูล\" onclick=\"editMemberDetails("+i+")\"></td></tr>";
                member += "</table>"
                if (i+1 < json.length) {
                    member += "<br>"
                }
            }
            $("#table_member_details").html(member);
        }
    }
    var data = JSON.stringify({"member_id":memberID});
    request.send(data);
}

function editMemberDetails(index){
    var id = parseInt($("#member_details_id_"+index).val());
    var memberNameTH = $("#member_name_th_id_"+index).val();
    var memberNameENG = $("#member_name_eng_id_"+index).val();
    var email = $("#email_id_"+index).val();
    var overtimeRate = parseFloat($("#overtime_rate_id_"+index).val());
    var ratePerDay = parseFloat($("#rate_per_day_id_"+index).val());
    var ratePerHour = parseFloat($("#rate_per_hour_id_"+index).val());
    var salary = parseFloat($("#salary_id_"+index).val());
    var incomeTax1 = parseFloat($("#income_tax_1_id_"+index).val());
    var socialSecurity = parseFloat($("#social_security_id_"+index).val());
    var incomeTax53Percentage = parseInt($("#income_tax_53_percentage_id_"+index).val());
    var status = $("#status_id_"+index).val();
    var travelExpense = parseFloat($("#travel_expense_id_"+index).val());

    var request = new XMLHttpRequest();
    var url = "/updateMemberDetails";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
        }
    }
    var data = JSON.stringify({"id":id,"member_name_th":memberNameTH,"member_name_eng":memberNameENG,
        "email":email,"overtime_rate":overtimeRate,"rate_per_day":ratePerDay,"rate_per_hour":ratePerHour,
        "salary":salary,"income_tax_1":incomeTax1,"social_security":socialSecurity,"income_tax_53_percentage":incomeTax53Percentage,
        "status":status,"travel_expense":travelExpense});    
    request.send(data);
}