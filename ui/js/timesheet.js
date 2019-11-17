function showSummary(){
    var date = $("#date_summary").val(); 
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;

    const monthNames = ["JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE","JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"];
      $(document).click(function(){
        $("#title_timesheet").text(month+"-"+monthNames[month-1]+year+"-TIMESHEET"); 
        $("#date_val").val(date); 
    });
    
    var request = new XMLHttpRequest();
    var url = "/showSummaryTimesheet";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) { 
            var json = JSON.parse(request.responseText);            
            var table = [];
            var count = [];
            var totalCoaching = [];
            var totalTraning = [];
            var totalOther = [];
            var totalIncomes = [];
            var totalSalary = [];
            var totalIncomeTax1 = [];
            var totalSocialSecurity = [];
            var totalNetSalary = [];
            var totalWage = [];
            var totalIncomeTax53 = [];
            var totalNetWage = [];
            var totalNetTransfer = [];
            var tableByCompany = "";
            var companyIndex;
            var companyName = [];

        if (json !== null){
            for (var index = 0; index < json.length; index++) {
                companyIndex = parseInt(json[index].company_id);
                count[companyIndex] = 0;
                table[companyIndex] = "";
                totalCoaching[companyIndex] = 0;
                totalTraning[companyIndex] = 0;
                totalOther[companyIndex] = 0;
                totalIncomes[companyIndex] = 0;
                totalSalary[companyIndex] = 0;
                totalIncomeTax1[companyIndex] = 0;
                totalSocialSecurity[companyIndex] = 0;
                totalNetSalary[companyIndex] = 0;
                totalWage[companyIndex] = 0;
                totalIncomeTax53[companyIndex] = 0;
                totalNetWage[companyIndex] = 0;
                totalNetTransfer[companyIndex] = 0;
                companyName[companyIndex] = ""
            }
            
            for (var index = 0; index < json.length; index++) {  
                companyIndex = parseInt(json[index].company_id);
                count[companyIndex]++;
                table[companyIndex] += "<tr id=\"row_summary_id_"+index+"\">";
                table[companyIndex] += "<td style=\"text-align: center;\">"+count[companyIndex]+"</td>";
                table[companyIndex] += "<td style=\"text-align: left;\" id=\"employee_name_eng_id_"+index+"\">"+json[index].employee_name_eng+"</td>";
                table[companyIndex] += "<td id=\"coaching_id_"+index+"\">"+setFormatMoney(json[index].coaching)+"</td>";
                totalCoaching[companyIndex] += json[index].coaching;
                table[companyIndex] += "<td id=\"training_id_"+index+"\">"+setFormatMoney(json[index].training)+"</td>";
                totalTraning[companyIndex] +=json[index].training;
                table[companyIndex] += "<td id=\"other_id_"+index+"\">"+setFormatMoney(json[index].other)+"</td>";
                totalOther[companyIndex] += json[index].other;
                table[companyIndex] += "<td class=\"yello\" id=\"total_incomes_id_"+index+"\">"+setFormatMoney(json[index].total_incomes)+"</td>";
                totalIncomes[companyIndex] += json[index].total_incomes;
                table[companyIndex] += "<td id=\"salary_id_"+index+"\">"+setFormatMoney(json[index].salary)+"</td>";
                totalSalary[companyIndex] += json[index].salary;
                table[companyIndex] += "<td id=\"income_tax_1_id_"+index+"\">"+setFormatMoney(json[index].income_tax_1)+"</td>";
                totalIncomeTax1[companyIndex] += json[index].income_tax_1;
                table[companyIndex] += "<td id=\"social_security_id_"+index+"\">"+setFormatMoney(json[index].social_security)+"</td>";
                totalSocialSecurity[companyIndex] += json[index].social_security;
                table[companyIndex] += "<td class=\"yello\" id=\"net_salary_id_"+index+"\">"+setFormatMoney(json[index].net_salary)+"</td>";
                totalNetSalary[companyIndex] += json[index].net_salary;
                table[companyIndex] += "<td id=\"wage_id_"+index+"\">"+setFormatMoney(json[index].wage)+"</td>";
                totalWage[companyIndex] += json[index].wage;
                table[companyIndex] += "<td id=\"income_tax_53_percentage_id_"+index+"\">"+json[index].income_tax_53_percentage+"&#37</td>";
                table[companyIndex] += "<td id=\"income_tax_53_id_"+index+"\">"+setFormatMoney(json[index].income_tax_53)+"</td>";
                totalIncomeTax53[companyIndex] += json[index].income_tax_53;
                table[companyIndex] += "<td id=\"net_wage_id_"+index+"\">"+setFormatMoney(json[index].net_wage)+"</td>";
                totalNetWage[companyIndex] +=json[index].net_wage;
                table[companyIndex] += "<td class=\"aqua\" id=\"net_transfer_id_"+index+"\">"+setFormatMoney(json[index].net_transfer)+"</td>";
                totalNetTransfer[companyIndex] += json[index].net_transfer;
                table[companyIndex] += "<td><select id=\"status_checking_transfer_"+index+"\"><option value=\""+json[index].status_checking_transfer+"\">"+json[index].status_checking_transfer+"</option>";
                table[companyIndex] += "<option value=\"รอการตรวจสอบ\">รอการตรวจสอบ</option>";
                table[companyIndex] += "<option value=\"โอนเงินเรียบร้อย\">โอนเงินเรียบร้อย</option>";
                table[companyIndex] += "<option value=\"ถูกต้อง\">ถูกต้อง</option>";
                table[companyIndex] += "<option value=\"ไม่ถูกต้อง\">ไม่ถูกต้อง</option>";
                table[companyIndex] += "</select></td>";
                table[companyIndex] += "<td><input type=\"text\" id=\"date_transfer_"+index+"\" value=\""+json[index].date_transfer+"\"></td>";
                table[companyIndex] += "<td><input type=\"text\" id=\"comment_"+index+"\" value=\""+json[index].comment+"\"></td>";
                table[companyIndex] += "<input type=\"hidden\" id=\"transaction_id_"+index+"\" value=\""+json[index].id+"\">";
                table[companyIndex] += "<input type=\"hidden\" id=\"employee_id_"+index+"\" value=\""+json[index].employee_id+"\">";
                table[companyIndex] += "<td>"+"<input class=\"button\" type=\"submit\" id=\"button_change_status_checking_transfer_id_"+index+"\" value=\"SUBMIT\" onclick=\"updateStatusTransfer("+index+")\"/>"+"</td>";
                table[companyIndex] += "</tr>";
                if (companyIndex == 1){
                    companyName[companyIndex] = "Siam Chamnankit"
                }else if (companyIndex == 2){
                    companyName[companyIndex] = "Shu ha ri"
                }else if (companyIndex == 3){
                    companyName[companyIndex] = "We Love Bug"
                }else{
                    companyName[companyIndex] = "Internship student AND Employee"
                    }
                }
            }
            
            for (var index = 0; index < 4; index++) {
                if (table[index]!=null){
                    tableByCompany += "<table border=\"1\" class=\"table_company\" width=\"2200\">";
                    tableByCompany += "<tr>"
                    tableByCompany += "<th rowspan=\"4\" class=\"blue\" style=\"width:30px;\">No</th>"
                    tableByCompany += "<tr><th rowspan=\"3\" id=\"company_name\" class=\"blue\" style=\"width:200px;\">"+companyName[index]
                    tableByCompany += "</th>"
                    tableByCompany += "<th colspan=\"4\" class=\"turquoise\"\">Income</th>"
                    tableByCompany += "<th colspan=\"4\" class=\"antiquewhite\"></th>"
                    tableByCompany += "<th colspan=\"4\" class=\"green\" >Wage Income of Withholding Income Tax (P.N.D.53)</th>"
                    tableByCompany += "<th class=\"blue\"></th>"
                    tableByCompany += "<th rowspan=\"3\" class=\"blue\">Inspection Status</th>"
                    tableByCompany += "<th rowspan=\"3\" class=\"blue\">Date For Transfer</th>"
                    tableByCompany += "<th rowspan=\"3\" class=\"blue\">Comment</th>"
                    tableByCompany += "<th rowspan=\"3\" class=\"blue\"></th>"
                    tableByCompany += "</tr>"
                    tableByCompany += "<tr>"
                    if (companyName[index] != "Shu ha ri"){
                        tableByCompany += "<td class=\"turquoise\"\" style=\"text-align: center;\">Coaching</td>"
                    }else{
                        tableByCompany += "<td class=\"turquoise\"\" style=\"text-align: center;\">Wage</td>"
                    }
                    tableByCompany += "<td class=\"turquoise\"\" style=\"text-align: center;\">Training</td>"
                    tableByCompany += "<td class=\"turquoise\"\" style=\"text-align: center;\">Other</td>"
                    tableByCompany += "<td class=\"turquoise\" style=\"text-align: center;\">Total Amount</td>"
                    tableByCompany += "<td class=\"antiquewhite\" style=\"text-align: center;\">Salary</td>"
                    tableByCompany += "<td class=\"antiquewhite\" style=\"text-align: center;\">Withholding Income Tax P.N.D.1</td>"
                    tableByCompany += "<td class=\"antiquewhite\" style=\"text-align: center;\">Social Security</td>"
                    tableByCompany += "<td class=\"antiquewhite\" style=\"text-align: center;\">Net Salary</td>"
                    tableByCompany += "<td class=\"green\" style=\"text-align: center;\">Wage</td>"
                    tableByCompany += "<td rowspan=\"2\" class=\"green\" style=\"text-align: center;\">Withholding Income Tax Rate P.N.D.53</td>"
                    tableByCompany += "<td class=\"green\" style=\"text-align: center;\">Withholding Income Tax P.N.D.53</td>"
                    tableByCompany += "<td class=\"green\" style=\"text-align: center;\">Net Wage</td>"
                    tableByCompany += "<td class=\"blue\" style=\"text-align: center;\">Net Transfer Amount</td>"
                    tableByCompany +=  "</tr>"                
                    tableByCompany +=  "<tr>"
                    tableByCompany +=  "<td id=\"total_coaching\" class=\"turquoise\"\">"+setFormatMoney(totalCoaching[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_traning\" class=\"turquoise\"\">"+setFormatMoney(totalTraning[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_other\" class=\"turquoise\"\">"+setFormatMoney(totalOther[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_incomes\" class=\"turquoise\"\">"+setFormatMoney(totalIncomes[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_salary\" class=\"antiquewhite\">"+setFormatMoney(totalSalary[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_income_tax_1\" class=\"antiquewhite\">"+setFormatMoney(totalIncomeTax1[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_social_security\" class=\"antiquewhite\">"+setFormatMoney(totalSocialSecurity[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_net_salary\" class=\"antiquewhite\">"+setFormatMoney(totalNetSalary[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_wage\" class=\"green\">"+setFormatMoney(totalWage[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_income_tax_53\" class=\"green\">"+setFormatMoney(totalIncomeTax53[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_net_wage\" class=\"green\"> "+setFormatMoney(totalNetWage[index])+"</td>"
                    tableByCompany +=  "<td id=\"total_net_transfer\" class=\"blue\">"+setFormatMoney(totalNetTransfer[index])+"</td>"
                    tableByCompany +=  "</tr>"
                    tableByCompany +=  table[index];
                    tableByCompany +=  "</table><br><br>"
                }
            }
            $("#table_company").html(tableByCompany);
        }
    }; 

    window.onscroll = function (){
        if (document.body.scrollTop > 10 || document.documentElement.scrollTop > 10) {
            document.getElementById("button_to_top").style.display="block";
        } else {
            document.getElementById("button_to_top").style.display="none";
        }
    };
    var data = JSON.stringify({"year":year, "month": month});
    request.send(data);
}

function setFormatMoney(amount){
    return "฿ "+ amount.toFixed(2).replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
}

function setFormatMoneyNoBath(amount){
    return amount.toFixed(2).replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
}

function updateStatusTransfer(index){
    var transactionID = $("#transaction_id_"+index).val();
    var statusTransfer = $("#status_checking_transfer_"+index).val();
    var dateTransfer = $("#date_transfer_"+index).val();
    var employeeID = $("#employee_id_"+index).val();
    var comment = $("#comment_"+index).val();
    var request = new XMLHttpRequest();
    var url = "/updateStatusCheckingTransfer";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token"));
    request.onreadystatechange = function () {
        if (request.status === 401){
            logout();
            deleteOauthState()
        }
    }
    var data = JSON.stringify({"employee_id":employeeID,"transaction_id":transactionID,"status":statusTransfer,"date":dateTransfer,"comment":comment});
    request.send(data); 
    window.location.replace(window.location.href)    
}

function createIncomeThenCalculatePayment(){
    new Promise(function(resolve, reject) {
        resolve(createIncome());
      }).then(calculatePayment());
}

function deleteIncomeThenCalculatePayment(index){
    new Promise(function(resolve, reject) {
        resolve(deleteIncome(index));
      }).then(calculatePayment());
}

function createIncome(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    employeeID = params.get("id");

    var full = $("#day").val()
    var year = parseInt(full.split("-")[0]);
    var month = parseInt(full.split("-")[1]);
    var day = parseInt(full.split("-")[2]);
    
    var fullStartTimeAm = $("#start_time_am").val();
    var startTimeAm = new Date("January 02, 2006 "+fullStartTimeAm+" UTC");
    
    var fullEndTimeAm = $("#end_time_am").val();
    var endTimeAm = new Date("January 02, 2006 "+fullEndTimeAm+" UTC");

    var fullStartTimePm = $("#start_time_pm").val();
    var startTimePm = new Date("January 02, 2006 "+fullStartTimePm+" UTC");

    var fullEndTimePm = $("#end_time_pm").val();
    var endTimePm = new Date("January 02, 2006 "+fullEndTimePm+" UTC");

    var coachingCustomerCharging = parseFloat($("#coaching_customer_charging").val());

    var coachingPaymentRate = parseFloat($("#coaching_payment_rate").val());
    
    var trainingWage = parseFloat($("#training_wage").val());

    var otherWage = parseFloat($("#other_wage").val());

    var companyID = parseInt($("#company_id").val());

    var description = $("#description").val();

    var request = new XMLHttpRequest();
    var url = "/createIncome";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token")); 
    request.onreadystatechange = function () {
        if (request.status === 401){
            alert("Your session has been expired, please log in again.")
            logout();
            if (deleteOauthState()){
                window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en");
            };
        }else if (request.status === 500){
            alert("Please fill out the information.")
            return
        }
    }
    var data = JSON.stringify({"year":year,"month":month,"employee_id":employeeID,"incomes":{"day":day,"start_time_am":startTimeAm,"end_time_am":endTimeAm,"start_time_pm":startTimePm,"end_time_pm":endTimePm,"coaching_customer_charging":coachingCustomerCharging,"coaching_payment_rate":coachingPaymentRate,"training_wage":trainingWage,"other_wage":otherWage,"company_id":companyID,"description":description}});
    request.send(data); 
    calculatePayment()
}

function calculatePayment() {
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    employeeID = params.get("id");
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;

    var request = new XMLHttpRequest();
    var url = "/calculatePayment";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token")); 
    request.onreadystatechange = function () {
        if (request.status === 401){
            logout();
            deleteOauthState();
        }
    }
    var data = JSON.stringify({"employee_id":employeeID,"year":year,"month":month});
    request.send(data)
    window.location.replace(window.location.href)
}

function goToSummaryTimesheetByEmployeeID(employeeID){
    $(document).ready(function(){
        location.href = "/home/showsummarybyid.html?date="+$("#date_val").val()+"&id="+employeeID;
    });
}

function showSummaryByID() {
    setCurrentDate();
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    employeeID = params.get("id");
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;
    var firstDay = new Date(fullDate.getFullYear(), fullDate.getMonth(), 1);
    var lastDay = new Date(fullDate.getFullYear(), fullDate.getMonth() + 1, 0);

    var src = "https://calendar.google.com/calendar/embed?height=600&amp;wkst=1&amp;bgcolor=%23ffffff&amp;ctz=Asia%2FBangkok&amp;src=ZW4udGgjaG9saWRheUBncm91cC52LmNhbGVuZGFyLmdvb2dsZS5jb20&amp;src=dGgudGgjaG9saWRheUBncm91cC52LmNhbGVuZGFyLmdvb2dsZS5jb20&amp;color=%230B8043&amp;color=%230B8043&amp;showTz=0&amp;showPrint=0&amp;showCalendars=0&amp;showTabs=0&amp;showNav=0&amp;dates=";
    var startDate = year.toString()+month.toString()+("0" + firstDay.getDate()).slice(-2)
    var endDate = year.toString()+month.toString()+("0" + lastDay.getDate()).slice(-2)
    
    var googleCalendarURL = "<iframe src=\""+src+startDate+"/"+endDate+"\" style=\"border-width:0\" width=\"600\" height=\"400\" frameborder=\"0\" scrolling=\"no\"></iframe>";
    
    if (date == null || employeeID == null) {
        alert("Please fill out the information.");
        location.href = document.referrer
    }

    const monthNames = ["January", "February", "March", "April", "May", "June","July", "August", "September", "October", "November", "December"];
    const monthNamesCapital = ["JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE","JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"];
    $(document).ready(function(){
      $("#title_timesheet_by_id").text(month+"-"+monthNamesCapital[month-1]+year+"-TIMESHEET");
      $("#month_name").html(monthNames[month-1]);
      $("#date_val").val(date);
    });

    var request = new XMLHttpRequest();
    var url = "/showTimesheetByEmployeeID";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var employeeNameENG = json.employee_name_eng;
		    var email = json.email;
		    var ratePerDay = json.rate_per_day;
		    var ratePerHour = json.rate_per_hour;
		    var totalHours = json.total_hours;
		    var totalCoachingCustomerCharging = json.total_coaching_customer_charging;
            var totalCoachingPaymentRate = json.total_coaching_payment_rate;
            var totalTrainigWage = json.total_training_wage;
            var totalOtherWage = json.total_other_wage;
            var paymentWage = json.payment_wage;
            var incomeList = "";
            
            if (json.incomes !== null) {
                for (var i = 0; i < json.incomes.length; i++) {
                    incomeList += "<tr id=\"income_company_"+json.incomes[i].company_id+"\">";
                    incomeList += "<td class=\"dark_green\">"+json.incomes[i].day+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].start_time_am)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].end_time_am)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].start_time_pm)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].end_time_pm)+"</td>";
                    incomeList += "<td>"+convertTimestampToTime(json.incomes[i].total_hours)+"</td>";
                    incomeList += "<td>"+setFormatMoney(json.incomes[i].coaching_customer_charging)+"</td>";
                    incomeList += "<td>"+setFormatMoney(json.incomes[i].coaching_payment_rate)+"</td>";
                    incomeList += "<td>"+setFormatMoney(json.incomes[i].training_wage)+"</td>";
                    incomeList += "<td>"+setFormatMoney(json.incomes[i].other_wage)+"</td>";
                    incomeList += "<td style=\"text-align: left;\">"+json.incomes[i].description+"</td>";
                    incomeList += "<td><input type=\"hidden\" id=\"income_id_"+i+"\" value=\""+json.incomes[i].id+"\">"
                    incomeList += "<input type=\"hidden\" id=\"employee_id_"+i+"\" value=\""+json.incomes[i].employee_id+"\">"
                    incomeList += "<input id=\"button_delete\" type=\"submit\" value=\"DELETE\" onclick=\"deleteIncomeThenCalculatePayment("+i+")\"/>"+"</td>"; 
                    setDateInIncomeFormat(json.incomes[i].day+1)
                    incomeList += "</tr>";
                }
                $("#table_timesheet").html(incomeList);
            }
            $("#employee_name_eng").html(employeeNameENG);
            $("#rate_per_day").html(setFormatMoneyNoBath(ratePerDay));
            $("#rate_per_hour").html(ratePerHour);
            $("#month_number").html(month);
            $("#year").html(year);

            $("#email").html(email);
            $("#thours").html(totalHours);
            $("#total_coaching_customer_charging").html(setFormatMoney(totalCoachingCustomerCharging));
            $("#total_coaching_payment_rate").html(setFormatMoney(totalCoachingPaymentRate));
            $("#total_trainig_wage").html(setFormatMoney(totalTrainigWage)); 
            $("#total_other_wage").html(setFormatMoney(totalOtherWage)); 
            $("#payment_wage").html(setFormatMoney(paymentWage));             
            $("#google_calendar").html(googleCalendarURL); 
            
        }
    }

    window.onscroll = function (){
        if (document.body.scrollTop > 10 || document.documentElement.scrollTop > 10) {
            document.getElementById("button_to_top").style.display="block";
        } else {
            document.getElementById("button_to_top").style.display="none";
        }
    };
    
    var data = JSON.stringify({"employee_id":employeeID,"year":year,"month":month});
    request.send(data); 

}

function setDateInIncomeFormat(lastIndex) {
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    date = params.get("date");
    var fullDate = new Date(date);
    var year = fullDate.getFullYear();
    var month = fullDate.getMonth()+1;
    if (lastIndex<10){
        $("#day").val(year+"-"+month+"-0"+lastIndex); 
    }else if (lastIndex>=10&&lastIndex<=31){
        $("#day").val(year+"-"+month+"-"+lastIndex); 
    }else{
        $("#day").val(year+"-"+month+"-01"); 
    }

}

function convertTimestampToTime(timestamp){
    var date = new Date(timestamp);
    datetext = date.toUTCString();
    datetext = datetext.split(' ')[4];
    return datetext
}

function deleteIncome(index){
    var incomeID = parseInt($("#income_id_"+index).val());    
    var employeeID = $("#employee_id_"+index).val()  
    var request = new XMLHttpRequest();
    var url = "/deleteIncomeItem";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token")); 
    request.onreadystatechange = function () {
        if (request.status === 401){
            logout();
            deleteOauthState()
        }
    }
    var data = JSON.stringify({"id":incomeID,"employee_id":employeeID});
    request.send(data);
    calculatePayment()
}

function goToProfilePage(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    var employeeID =  params.get("id");
    if (employeeID != null){
        location.href = "/home/employeedetails.html?id="+employeeID;
    }else{
        alert("Please fill out the information.")
    }
}

function goToSummaryByIDBackPage(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    var employeeID =  params.get("id");

    var currentTime = new Date();
    var currentYear = String(currentTime.getFullYear());
    var currentMonth = String(currentTime.getMonth()+1);
    var today = currentYear + "-" + currentMonth;
    if (employeeID != null){
        location.href = "/home/showsummarybyid.html?date="+today+"&id="+employeeID;
    }
}

function goToPaymentSummaryAnnualPage(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    var employeeID =  params.get("id");
    $("#date_val").val(date); 
    location.href = "/home/paymentsummaryannual.html?id="+employeeID+"&year="+$("#date_val").val().split("-")[0];
}

function paymentSummaryAnnual(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    var employeeID = params.get("id");
    var year = parseInt(params.get("year"));

    var request = new XMLHttpRequest();
    var url = "/showSummaryInYear";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var table = "";
            const monthNames = ["January", "February", "March", "April", "May", "June","July", "August", "September", "October", "November", "December"];
            for (var index = 0; index < json.transaction_timesheets.length; index++) {  
                table += "<tr id=\"row_summary_in_year_id_"+index+"\">";
                table += "<td style=\"text-align: left;\" id=\"month_name_id_"+index+"\">"+monthNames[json.transaction_timesheets[index].month-1]+"</td>";
                if (json.transaction_timesheets[index].company_id == 1){
                    table += "<td style=\"text-align: left;\" id=\"company_name_"+index+"\">Siam Chamnankit</td>";
                }else if (json.transaction_timesheets[index].company_id == 2){
                    table += "<td style=\"text-align: left;\" id=\"company_name_"+index+"\">Shu Ha Ri</td>";
                }else {
                    table += "<td style=\"text-align: left;\" id=\"company_name_"+index+"\">We Love Bug</td>";
                }
                table += "<td id=\"coaching_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].coaching)+"</td>";
                table += "<td id=\"training_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].training)+"</td>";
                table += "<td id=\"other_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].other)+"</td>";
                table += "<td class=\"yello\" id=\"total_incomes_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].total_incomes)+"</td>";
                table += "<td id=\"salary_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].salary)+"</td>";
                table += "<td id=\"income_tax_1_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].income_tax_1)+"</td>";
                table += "<td id=\"social_security_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].social_security)+"</td>";
                table += "<td class=\"yello\" id=\"net_salary_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].net_salary)+"</td>";
                table += "<td id=\"wage_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].wage)+"</td>";
                table += "<td id=\"income_tax_53_percentage_id_"+index+"\">"+json.transaction_timesheets[index].income_tax_53_percentage+"&#37</td>";
                table += "<td id=\"income_tax_53_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].income_tax_53)+"</td>";
                table += "<td id=\"net_wage_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].net_wage)+"</td>";
                table += "<td class=\"aqua\" id=\"net_transfer_id_"+index+"\">"+setFormatMoney(json.transaction_timesheets[index].net_transfer)+"</td>";
                table += "<td id=\"status_checking_transfer_"+index+"\">"+json.transaction_timesheets[index].status_checking_transfer+"</td>";
                table += "<td id=\"date_transfer_"+index+"\">"+json.transaction_timesheets[index].date_transfer+"</td>";
                table += "<td id=\"comment_"+index+"\">"+json.transaction_timesheets[index].comment+"</td></tr>";
            }
            
            $("#employee_name_eng").html(json.transaction_timesheets[0].employee_name_eng); 
            $("#payment_year_summary_annual").html(json.year); 
            $("#table_summary_in_year").html(table); 
            $("#total_coaching_in_year").html(setFormatMoney(json.total_coaching_in_year)); 
            $("#total_income_tax_1_in_year").html(setFormatMoney(json.total_income_tax_1_in_year)); 
            $("#total_income_tax_53_in_year").html(setFormatMoney(json.total_income_tax_53_in_year)); 
            $("#total_incomes_in_year").html(setFormatMoney(json.total_incomes_in_year)); 
            $("#total_net_salary_in_year").html(setFormatMoney(json.total_net_salary_in_year)); 
            $("#total_net_transfer_in_year").html(setFormatMoney(json.total_net_transfer_in_year)); 
            $("#total_net_wage_in_year").html(setFormatMoney(json.total_net_wage_in_year)); 
            $("#total_other_in_year").html(setFormatMoney(json.total_other_in_year)); 
            $("#total_salary_in_year").html(setFormatMoney(json.total_salary_in_year)); 
            $("#total_social_security_in_year").html(setFormatMoney(json.total_social_security_in_year)); 
            $("#total_training_in_year").html(setFormatMoney(json.total_training_in_year)); 
            $("#total_wage_in_year").html(setFormatMoney(json.total_wage_in_year));  
        }
    }   
    var data = JSON.stringify({"employee_id":employeeID,"year":year});
    request.send(data);
}

function getEmployeeByID(){
    var urlString = window.location.href
    var url = new URL(urlString);
    var params = new URLSearchParams(url.search);
    employeeID = params.get("id");

    var request = new XMLHttpRequest();
    var url = "/showEmployeeDetailsByEmployeeID";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var employee = "";
            for (var i = 0; i < json.length; i++) {
                employee += "<table width=\"600\">"
                employee += "<tr><th>Company Name</th><td id=\"company_name_"+i+"\">"+json[i].company_name+"</td></tr>";
                employee += "<tr><th>Name (Thai)</th><td id=\"employee_name_th_id_"+i+"\">"+json[i].employee_name_th+"</td></tr>";
                employee += "<tr><th>Name (English)</th><td id=\"employee_name_eng_id_"+i+"\">"+json[i].employee_name_eng+"</td></tr>";
                employee += "<tr><th>E-mail</th><td id=\"email_id_"+i+"\">"+json[i].email+"</td></tr>";
                employee += "<tr><th>Rate Per Day (THB)</th><td><input type=\"number\" id=\"rate_per_day_id_"+i+"\" value=\""+json[i].rate_per_day.toFixed(2)+"\"></td></tr>";
                employee += "<tr><th>Rate Per Hour (THB)</th><td><input type=\"number\" id=\"rate_per_hour_id_"+i+"\" value=\""+json[i].rate_per_hour+"\"></td></tr>";
                employee += "<tr><th>Salary (THB)</th><td><input type=\"number\" id=\"salary_id_"+i+"\" value=\""+json[i].salary.toFixed(2)+"\"></td></tr>";
                employee += "<tr><th>Withholding Income Tax P.N.D.1 (THB)</th><td><input type=\"number\" id=\"income_tax_1_id_"+i+"\" value=\""+json[i].income_tax_1.toFixed(2)+"\"></td></tr>";
                employee += "<tr><th>Social Security Tax</th><td><input type=\"number\" id=\"social_security_id_"+i+"\" value=\""+json[i].social_security.toFixed(2)+"\"></td></tr>";
                employee += "<tr><th>Withholding Income Tax Rate P.N.D.53 (Percentage)</th><td><input type=\"number\" id=\"income_tax_53_percentage_id_"+i+"\" value=\""+json[i].income_tax_53_percentage+"\"></td></tr>";
                employee += "<tr><th>Type of Income</th><td><select id=\"status_id_"+i+"\">";
                if (json[i].status == "wage"){
                    employee += "<option value=\""+json[i].status+"\">ค่าจ้างรายวัน (wage)</option>";
                    employee += "<option value=\"salary\">เงินเดือน (salary)</option>";
                }else{
                    employee += "<option value=\""+json[i].status+"\">เงินเดือน (salary)</option>";
                    employee += "<option value=\"wage\">ค่าจ้างรายวัน (wage)</option>";
                }
                employee += "</select></td></tr>";
                employee += "<tr><th>Travel Expenses (THB)</th><td><input type=\"number\" id=\"travel_expense_id_"+i+"\" value=\""+json[i].travel_expense.toFixed(2)+"\"></td></tr>";
                employee += "<input type=\"hidden\" id=\"employee_details_id_"+i+"\" value=\""+json[i].id+"\">";
                employee += "<input type=\"hidden\" id=\"employee_id_"+i+"\" value=\""+employeeID+"\">";
                employee += "<input type=\"hidden\" d=\"company_id_"+i+"\" value=\""+json[i].company_id+"\">";
                employee += "<input type=\"hidden\" d=\"company_id_"+i+"\" value=\""+json[i].company_id+"\">";
                employee += "<input type=\"hidden\" d=\"company_id_"+i+"\" value=\""+json[i].company_id+"\">";
                employee += "<tr><td></td><td><input class=\"btn btn-primary mb-2\" type=\"submit\" id=\"button_edit_employee_id_"+i+"\" value=\"EDIT\" onclick=\"editEmployeeDetails("+i+")\"></td></tr>";                                    
                employee += "</table>"
                if (i+1 < json.length) {
                    employee += "<br><br><br>"
                }
            }
            
            $("#table_employee_details").html(employee);
        }
    }

    window.onscroll = function (){
        if (document.body.scrollTop > 10 || document.documentElement.scrollTop > 10) {
            document.getElementById("button_to_top").style.display="block";
        } else {
            document.getElementById("button_to_top").style.display="none";
        }
    };
    
    var data = JSON.stringify({"employee_id":employeeID});
    request.send(data);
}

function editEmployeeDetails(index){    
    var id = parseInt($("#employee_details_id_"+index).val());
    var employeeID = $("#employee_id_"+index).val();
    var employeeNameTH = $("#employee_name_th_id_"+index).text();
    var employeeNameENG = $("#employee_name_eng_id_"+index).text();
    var email = $("#email_id_"+index).text();
    var ratePerDay = parseFloat($("#rate_per_day_id_"+index).val());
    var ratePerHour = parseFloat($("#rate_per_hour_id_"+index).val());
    var salary = parseFloat($("#salary_id_"+index).val());
    var incomeTax1 = parseFloat($("#income_tax_1_id_"+index).val());
    var socialSecurity = parseFloat($("#social_security_id_"+index).val());
    var incomeTax53Percentage = parseInt($("#income_tax_53_percentage_id_"+index).val());
    var status = $("#status_id_"+index).val();
    var travelExpense = parseFloat($("#travel_expense_id_"+index).val());

    var request = new XMLHttpRequest();
    var url = "/updateEmployeeDetails";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token")); 
    request.onreadystatechange = function () {
        if (request.status === 401){
            logout();
            deleteOauthState()
        }
    }
    var data = JSON.stringify({"id":id,"employee_id":employeeID,"employee_name_th":employeeNameTH,"employee_name_eng":employeeNameENG,
        "email":email,"rate_per_day":ratePerDay,"rate_per_hour":ratePerHour,
        "salary":salary,"income_tax_1":incomeTax1,"social_security":socialSecurity,"income_tax_53_percentage":incomeTax53Percentage,
        "status":status,"travel_expense":travelExpense}); 
    
    request.send(data);
    window.location.replace(window.location.href)
}

function setCurrentDate(){
    var currentTime = new Date();
    var currentYear = String(currentTime.getFullYear());
    var currentMonth = String(currentTime.getMonth()+1);
    var today = currentYear + "-" + currentMonth;
    $(document).ready(function(){
        $("#date_summary").val(today);  
        $("#date_val").val(today);
        $("#date").val(today); 
        setInitialHome();
        if (window.location.pathname === "/home/"){
            showSummary();
        }
    });

    const monthNames = ["JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE","JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"];
    $(document).ready(function(){
        $("#title_timesheet").text(currentMonth+"-"+monthNames[currentMonth-1]+currentYear+"-TIMESHEET");  
    }); 
}

function showProfile(){
    var request = new XMLHttpRequest();
    var url = "/showProfile";
    request.open("GET", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization",getCookie("id_token"));
    request.onreadystatechange = function () {
        if (request.status === 401){
            logout();
            if(deleteOauthState()){
                window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
            };
        }
        if (request.readyState === 4 && request.status === 200) {
            var json = JSON.parse(request.responseText);
            var picture = "<img class=\"circular--square\" src=\""+json.picture+"\">"
            $("#picture_profile").html(picture);
            $("#email_profile").html(json.email);  
        }
    }   
    request.send();
}

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
}

function setCookie(cname, cvalue, exdays) {
    var date = new Date();
    date.setTime(date.getTime() + (exdays*24*60*60*1000));
    var expires = "expires="+ date.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function setInitialHome(){    
    var loginButton = "<a id=\"button_login\" href=\"/login\"><span class=\"glyphicon glyphicon-log-in\"></span> Login</a>"
    var logoutButton = "<a id=\"button_logout\"><span class=\"glyphicon glyphicon-log-in\"></span> Logout</a>"

    $(document).ready(function(){
        if (getCookie("id_token") == ""){
                window.location.replace("/login")
        }
        if (getCookie("oauthstate") != ""){
            $("#login").html(logoutButton);
            showProfile();
            
            $("#button_logout").click(function(){
                logout();
                if (deleteOauthState()){
                    window.location.replace("https://mail.google.com/mail/u/0/?logout&hl=en")
                };
            });
        }else{  
            $("#login").html(loginButton);
        } 
    });
}

function logout(){
    var request = new XMLHttpRequest();
    var url = "/logout";
    request.open("POST", url, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", getCookie("id_token"));
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {       
        }
    }   
    request.send();
}

function deleteOauthState(){
    var request = new XMLHttpRequest();
    var url = "/deleteOauthState";
    request.open("GET", url, true);
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {    
        }
    }   
    request.send();
    return true
}

function topFunction() {
    document.body.scrollTop = 0;
    document.documentElement.scrollTop = 0;
}