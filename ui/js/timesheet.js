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
                    siamChamnankit += "<tr>";
                    siamChamnankit += "<td>"+countSiamChamnankit+"</td>";
                    siamChamnankit += "<td>"+json[i-1].member_name_th+"</td>";
                    siamChamnankit += "<td>"+json[i-1].coaching+"</td>";
                    totalCoachingSiamChamnankit += json[i-1].coaching;
                    siamChamnankit += "<td>"+json[i-1].training+"</td>";
                    totalTraningSiamChamnankit +=json[i-1].training;
                    siamChamnankit += "<td>"+json[i-1].other+"</td>";
                    totalOtherSiamChamnankit += json[i-1].other;
                    siamChamnankit += "<td>"+json[i-1].total_incomes+"</td>";
                    totalIncomesSiamChamnankit += json[i-1].total_incomes;
                    siamChamnankit += "<td>"+json[i-1].salary+"</td>";
                    totalSalarySiamChamnankit += json[i-1].salary;
                    siamChamnankit += "<td>"+json[i-1].income_tax_1+"</td>";
                    totalIncomeTax1SiamChamnankit += json[i-1].income_tax_1;
                    siamChamnankit += "<td>"+json[i-1].social_security+"</td>";
                    totalSocialSecuritySiamChamnankit += json[i-1].social_security;
                    siamChamnankit += "<td>"+json[i-1].net_salary+"</td>";
                    totalNetSalarySiamChamnankit += json[i-1].net_salary;
                    siamChamnankit += "<td>"+json[i-1].wage+"</td>";
                    totalWageSiamChamnankit += json[i-1].wage;
                    siamChamnankit += "<td>"+json[i-1].income_tax_53_percentage+"</td>";
                    siamChamnankit += "<td>"+json[i-1].income_tax_53+"</td>";
                    totalIncomeTax53SiamChamnankit += json[i-1].income_tax_53;
                    siamChamnankit += "<td>"+json[i-1].net_wage+"</td>";
                    totalNetWageSiamChamnankit +=json[i-1].net_wage;
                    siamChamnankit += "<td>"+json[i-1].net_transfer+"</td>";
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
                    siamChamnankit += "<td>"+"<input type=\"submit\" value=\"เปลี่ยนสถานะ\" onclick=\"updateStatusTransfer("+i+")\"/>"+"</td>";
                    siamChamnankit += "</tr>";
            console.log(json[i-1].company);
            
                }else{
                    countShuhari++;
                    shuhari += "<tr>";
                    shuhari += "<td>"+countShuhari+"</td>";
                    shuhari += "<td>"+json[i-1].member_name_th+"</td>";
                    shuhari += "<td>"+json[i-1].coaching+"</td>";
                    totalCoachingShuhari += json[i-1].coaching;
                    shuhari += "<td>"+json[i-1].training+"</td>";
                    totalTraningShuhari += json[i-1].training;
                    shuhari += "<td>"+json[i-1].other+"</td>";
                    totalOtherShuhari += json[i-1].other;
                    shuhari += "<td>"+json[i-1].total_incomes+"</td>";
                    totalIncomesShuhari += json[i-1].total_incomes;
                    shuhari += "<td>"+json[i-1].salary+"</td>";
                    totalSalaryShuhari += json[i-1].salary;
                    shuhari += "<td>"+json[i-1].income_tax_1+"</td>";
                    totalIncomeTax1Shuhari += json[i-1].income_tax_1;
                    shuhari += "<td>"+json[i-1].social_security+"</td>";
                    totalSocialSecurityShuhari += json[i-1].social_security;
                    shuhari += "<td>"+json[i-1].net_salary+"</td>";
                    totalNetSalaryShuhari += json[i-1].net_salary;
                    shuhari += "<td>"+json[i-1].wage+"</td>";
                    totalWageShuhari += json[i-1].wage;
                    shuhari += "<td>"+json[i-1].income_tax_53_percentage+"</td>";
                    shuhari += "<td>"+json[i-1].income_tax_53+"</td>";
                    totalIncomeTax53Shuhari += json[i-1].income_tax_53;
                    shuhari += "<td>"+json[i-1].net_wage+"</td>";
                    totalNetWageShuhari += json[i-1].net_wage;
                    shuhari += "<td>"+json[i-1].net_transfer+"</td>";
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
                    shuhari += "<td>"+"<input type=\"submit\" value=\"เปลี่ยนสถานะ\" onclick=\"updateStatusTransfer("+i+")\"/>"+"</td>";
                    shuhari += "</tr>";
                    console.log(json[i-1].company);
            
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
    
    addIncomeToTimesheet(memberID,year,month)
    calculatePayment(memberID,year,month)
    window.location.replace(window.location.href)    
}

function addIncomeToTimesheet(memberID,year,month){
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
}

function calculatePayment(memberID,year,month) {
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
            console.log(totalCoachingPaymentRate);
            
            if (json.incomes !== null) {
                for (var i = 0; i < json.incomes.length; i++) {
                    incomeList += "<tr>";
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
                    // incomeList += "<td>"+json.incomes[i].company+"</td>";
                    incomeList += "<td>"+json.incomes[i].description+"</td>";
                    incomeList += "</tr>";  
                }
                $("#table_timesheet").html(incomeList);
            }
            
            $("#member_name_eng").html(memberNameENG);
            $("#email").html(email);
            $("#overtime_rate").html(overtimeRate);
            $("#ratePerDay").html(ratePerDay);
            $("#ratePerHour").html(ratePerHour);
            $("#month").html(month);
            $("#monthString").html(monthString);
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