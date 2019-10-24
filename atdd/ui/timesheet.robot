*** Settings ***
Library    SeleniumLibrary

***Variables***
${URL_PAYMENTS}    http://localhost:8080/
${URL_PAYMENTS_YEAR_MONTH}    http://localhost:8080/?date_summary=2018-12

*** Test Case ***
ดูผลสรุปในหน้า PAYMENTS
    เปิด Browser
    ใส่เดือนและปีที่ต้องการดูสรุปผล
    กดปุ่มยืนยัน
    เข้าสู่หน้าสรุปผลค่าจ้างเดือนและปีนั้น
    ต้องเจอพนักงานในตารางแรก
    ปิด Browser

***Keywords***
เปิด Browser 
    Open Browser    ${URL_PAYMENTS}    chrome

ใส่เดือนและปีที่ต้องการดูสรุปผล
    Input Text    id=date_summary    12\t2018

กดปุ่มยืนยัน
    Click Button    id=button_show_summary

เข้าสู่หน้าสรุปผลค่าจ้างเดือนและปีนั้น
    Location Should Be    ${URL_PAYMENTS_YEAR_MONTH}

ต้องเจอพนักงานในตารางแรก
    Wait Until Page Contains Element    row_summary_id_1
    Element Text Should Be    id=member_name_th_id_1    ประธาน ด่านสกุลเจริญกิจ
    Element Text Should Be    id=coaching_id_1    75000
    Element Text Should Be    id=training_id_1    30000
    Element Text Should Be    id=other_id_1    40000
    Element Text Should Be    id=total_incomes_id_1    145000
    Element Text Should Be    id=salary_id_1    80000
    Element Text Should Be    id=income_tax_1_id_1    5000
    Element Text Should Be    id=social_security_id_1    0
    Element Text Should Be    id=net_salary_id_1    75000
    Element Text Should Be    id=wage_id_1    65000
    Element Text Should Be    id=income_tax_53_percentage_id_1    10
    Element Text Should Be    id=income_tax_53_id_1    6500
    Element Text Should Be    id=net_wage_id_1    58500
    Element Text Should Be    id=net_transfer_id_1    133500
    Select From List By Value    id=status_checking_transfer_1    รอการตรวจสอบ

ปิด Browser
    Close Browser