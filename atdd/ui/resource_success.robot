*** Variables ***
${LOCAL_URL}    http://localhost:8080/login
${BROWSER}    chrome

*** Keywords ***
เข้าสู่ระบบ
    Maximize Browser Window
    Wait Until Page Contains    ลงชื่อเข้าใช้งาน
    Input Text    id=identifierId    nareenart@scrum123.com
    Click Element    id=identifierNext
    Wait Until Page Contains    ยินดีต้อนรับ
    Wait Until Element Is Visible    name=password
    Input Password    name=password    delight2538
    Click Element    id=passwordNext

ใส่เดือนและปีที่ต้องการดูสรุปผล
    [Arguments]    ${DATE}    ${TITLE}
    Wait Until Page Contains    PAYMENTS
    Input Text    id=date_summary    ${DATE}
    Click Element    id=button_show_summary
    Element Text Should Be    id=title_timesheet    ${TITLE}

แสดงรายละเอียดของพนักงาน ID
    [Arguments]    ${ID}    ${NAME}    ${COACHING}    ${TRAINING}    ${OTHER}    ${TOTAL_INCOME}    ${SALARY}    ${INCOME_TAX_1}    ${SOCIAL_SECURITY}    ${NET_SALARY}    ${WAGE}    ${INCOME_TAX_53_PERCENTAGE}    ${INCOME_TAX_53}    ${NET_WAGE}    ${NET_TRANSFER}    ${STATUS_CHECKING_TRANSFER}    ${DATE}     ${COMMENT} 
    Wait Until Page Contains Element    row_summary_id_${ID}
    Element Text Should Be    id=employee_name_eng_id_${ID}    ${NAME}
    Element Text Should Be    id=coaching_id_${ID}    ${COACHING}
    Element Text Should Be    id=training_id_${ID}    ${TRAINING}
    Element Text Should Be    id=other_id_${ID}    ${OTHER}
    Element Text Should Be    id=total_incomes_id_${ID}    ${TOTAL_INCOME}
    Element Text Should Be    id=salary_id_${ID}    ${SALARY}
    Element Text Should Be    id=income_tax_1_id_${ID}    ${INCOME_TAX_1}
    Element Text Should Be    id=social_security_id_${ID}    ${SOCIAL_SECURITY}
    Element Text Should Be    id=net_salary_id_${ID}    ${NET_SALARY}
    Element Text Should Be    id=wage_id_${ID}    ${WAGE}
    Element Text Should Be    id=income_tax_53_percentage_id_${ID}    ${income_tax_53_percentage}
    Element Text Should Be    id=income_tax_53_id_${ID}    ${INCOME_TAX_53}
    Element Text Should Be    id=net_wage_id_${ID}    ${NET_WAGE}
    Element Text Should Be    id=net_transfer_id_${ID}    ${NET_TRANSFER}
    Select From List By Value    id=status_checking_transfer_${ID}    ${STATUS_CHECKING_TRANSFER}
    Textfield Should Contain    id=date_transfer_${ID}    ${DATE}
    Textfield Should Contain    id=comment_${ID}     ${COMMENT} 


เปลี่ยนสถานะการโอน วันที่โอน และหมายเหตุ
    [Arguments]    ${ID}    ${STATUS}    ${DATE}    ${COMMENT}
    Select From List By Label    id=status_checking_transfer_${ID}    ${STATUS}
    Input Text    id=date_transfer_${ID}    ${DATE}
    Input Text    id=comment_${ID}    ${COMMENT}
    Click Button    id=button_change_status_checking_transfer_id_${ID}

เลือกแสดงรายละเอียดสรุปผลรายบุคคล
    [Arguments]    ${ID}
    Click Element    id=${ID}
    
แสดงข้อมูลพนักงานในหน้าสรุปผลค่าจ้างรายบุคคลของเดือนและปี
    [Arguments]    ${TITLE}    ${NAME}    ${EMAIL}    ${RATE_PER_DAY}    ${RATE_PER_HOUR}    ${MONTH_NUMBER}    ${MONTH_NAME}    ${YEAR}
    Element Text Should Be    id=title_timesheet_by_id    ${TITLE}
    Element Text Should Be    id=employee_name_eng    ${NAME}
    Element Text Should Be    id=email    ${EMAIL}
    Element Text Should Be    id=rate_per_day    ${RATE_PER_DAY}
    Element Text Should Be    id=rate_per_hour    ${RATE_PER_HOUR}
    Element Text Should Be    id=month_number    ${MONTH_NUMBER}
    Element Text Should Be    id=month_name    ${MONTH_NAME}
    Element Text Should Be    id=year    ${YEAR}

แสดงสรุปผลค่าจ้างรายบุคคลทั้งหมดเป็น
    [Arguments]    ${TOTAL_HOURS}    ${PAYMENT_WAGE}    ${TOTAL_COACHING_CUSTOMER_CHARCHING}    ${TOTAL_COACHING_PAYMENT_RATE}    ${TOTAL_TRAINING_WAGE}    ${TOTAL_OTHER_WAGE}
    Element Text Should Be    id=thours    ${TOTAL_HOURS}
    Element Text Should Be    id=payment_wage    ${PAYMENT_WAGE}
    Element Text Should Be    id=total_coaching_customer_charging    ${TOTAL_COACHING_CUSTOMER_CHARCHING}
    Element Text Should Be    id=total_coaching_payment_rate    ${TOTAL_COACHING_PAYMENT_RATE}
    Element Text Should Be    id=total_trainig_wage    ${TOTAL_TRAINING_WAGE}
    Element Text Should Be    id=total_other_wage    ${TOTAL_OTHER_WAGE}

เพิ่มค่าจ้างรายวันในส่วนของ Coaching (THB)
    [Arguments]    ${DAY}    ${COACHING_CUSTOMER_CHARGING}    ${COACHING_PAYMENT_RATE}    ${COMPANY}    ${DESCRIPTION}
    Input Text    id=day    ${DAY}
    Select From List By Label    id=coaching_customer_charging    ${COACHING_CUSTOMER_CHARGING}
    Select From List By Label    id=coaching_payment_rate    ${COACHING_PAYMENT_RATE}
    Select From List By Label    id=company_id    ${COMPANY}
    Input Text    id=description    ${DESCRIPTION}
    Click Element    id=button_add_income_item


เพิ่มค่าจ้างรายวันในส่วนของ Training Wage (THB)
    [Arguments]    ${DAY}    ${TRAINING_WAGE}    ${COMPANY}    ${DESCRIPTION}
    Input Text    id=day    ${DAY}
    Select From List By Label    id=training_wage    ${TRAINING_WAGE}
    Select From List By Label    id=company_id    ${COMPANY}
    Input Text    id=description    ${DESCRIPTION}
    Click Element    id=button_add_income_item


เพิ่มค่าจ้างรายวันในส่วนของ Other Wage (THB)
    [Arguments]    ${DAY}    ${OTHER_WAGE}    ${COMPANY}    ${DESCRIPTION}
    Input Text    id=day    ${DAY}
    Select From List By Label    id=other_wage    ${OTHER_WAGE}
    Select From List By Label    id=company_id    ${COMPANY}
    Input Text    id=description    ${DESCRIPTION}
    Click Element    id=button_add_income_item

เข้าสู่หน้าแก้ไขข้อมูลพนักงาน
    [Arguments]    ${ID}
    Click Element    id=${ID}
    Click Element    id=button_edit_employee

แสดงหน้าแก้ไขข้อมูลพนักงาน
    [Arguments]    ${ID}    ${COMPANY}    ${์NAME_TH}    ${NAME_ENG}    ${EMAIL}    ${RATE_PER_DAY}    ${RATE_PER_HOUR}    ${SALARY}    ${INCOME_TAX_1}    ${SOCIAL_SECURITY}    ${INCOME_TAX_53_PERCENTAGE}    ${STATUS}    ${TRAVEL_EXPESE}
    
    Element Text Should Be    id=company_name_${ID}    ${COMPANY}
    Element Text Should Be    id=employee_name_th_id_${ID}    ${์NAME_TH}
    Element Text Should Be    id=employee_name_eng_id_${ID}    ${NAME_ENG}
    Element Text Should Be    id=email_id_${ID}    ${EMAIL}
    Textfield Should Contain    id=rate_per_day_id_${ID}    ${RATE_PER_DAY}
    Textfield Should Contain    id=rate_per_hour_id_${ID}    ${RATE_PER_HOUR}
    Textfield Should Contain    id=salary_id_${ID}    ${SALARY}
    Textfield Should Contain    id=income_tax_1_id_${ID}    ${INCOME_TAX_1}
    Textfield Should Contain    id=social_security_id_${ID}    ${SOCIAL_SECURITY}
    Textfield Should Contain    id=income_tax_53_percentage_id_${ID}    ${income_tax_53_percentage}
    Select From List By Value   id=status_id_${ID}    ${STATUS}
    Textfield Should Contain    id=travel_expense_id_${ID}    ${TRAVEL_EXPESE}

แก้ไขข้อมูลเงินเดือนพนักงาน
    [Arguments]    ${ID}    ${SALARY}
    Input Text    id=salary_id_${ID}    ${SALARY}
    Scroll Element Into View    id=button_edit_employee_id_${ID}
    Click Element    id=button_edit_employee_id_${ID}

ออกจากระบบ
    Click Element    id=button_logout