*** Settings ***
Library    SeleniumLibrary

***Variables***
${URL_PAYMENTS}    http://localhost:8080/

*** Test Case ***
ดูผลสรุปในหน้า PAYMENTS และเปลี่ยนสถานะการโอนเงินเป็น ถูกต้อง สำเร็จของ PRATHAN
    เปิด Browser
    ใส่เดือนและปีที่ต้องการดูสรุปผล
    กดปุ่มยืนยันดูสรุปผล    
    เข้าสู่หน้าสรุปผลค่าจ้างเดือนและปีนั้น
    ต้องเจอพนักงานในตารางแรก    1    ประธาน ด่านสกุลเจริญกิจ    75000    30000    40000    145000    80000    5000    0    75000    65000    10    6500    58500    133500    รอการตรวจสอบ
    เปลี่ยนสถานะการตรวจสอบ    1    ถูกต้อง
    ใส่วันที่โอนเงิน    1    28/12/2018
    ใส่หมายเหตุ    1     ค่าตั๋วที่ออกไปก่อน = 169,380.00 บาท
    กดปุ่มเปลี่ยนสถานะ    1
    ตรวจสอบสถานะการโอน    1    ถูกต้อง    28/12/2018     ค่าตั๋วที่ออกไปก่อน = 169,380.00 บาท
    ปิด Browser

ดูสรุปค่าจ้างรายบุคคลของ PRATHAN และสามารถเพิ่มค่าจ้างรายวันในวันที่ 28 สำเร็จ
    เปิด Browser
    ใส่เดือนและปีที่ต้องการดูสรุปผลรายบุคคล    12\t2018
    ใส่ชื่อที่ต้องการดูสรุปผลรายบุคคล     PRATHAN
    กดปุ่มยืนยันดูสรุปค่าจ้างรายบุคคล
    เข้าสู่หน้าสรุปผลค่าจ้างรายบุคคลของเดือนและปีนั้น    Prathan Dansakulcharoenkit    prathan@scrum123.com     0    15000    1875    12    December    2018
    ชั่วโมงการทำงานทั้งหมด    144:00:00
    ค่าจ้างทั้งหมดเป็น    185000    15000    75000    70000    40000
    ใส่วันที่ต้องการเพิ่มค่าจ้างรายวัน    28
    ใส่เวลาเริ่มงานช่วงเช้า    090000
    ใส่เวลาจบงานช่วงเช้า    120000
    ใส่เวลาเริ่มงานช่วงบ่าย    130000
    ใส่เวลาจบงานช่วงบ่าย    180000
    ใส่ชั่วโมงการทำงานล่วงเวลา    0
    ใส่ชั่วโมงการทำงานรวมของวัน    080000
    ใส่ค่า Coaching Customer Charging (THB)    ฿ 15,000.00
    ใส่ค่า Coaching Payment Rate (THB)    ฿ 10,000.00
    ใส่ค่า Training Wage (THB)    ฿ 0.00
    ใส่ค่า Other Wage (THB)    ฿ 0.00
    ใส่บริษัทที่เป็นผู้รับผิดชอบ    Shuhari
    ใส่คำอธิบายสถานที่หรือหมายเหตุ    Siam Chamnankit and SHR operation
    กดปุ่มยืนยันเพิ่มค่าจ้างรายวัน
    ปิด Browser

แก้ไขข้อมูลเงินเดือนเป็น 30000 ของ NAREENART สำเร็จ
    เปิด Browser
    ใส่เดือนและปีที่ต้องการดูสรุปผลรายบุคคล    12\t2019
    ใส่ชื่อที่ต้องการดูสรุปผลรายบุคคล     NAREENART
    กดปุ่มยืนยันดูสรุปค่าจ้างรายบุคคล
    กดปุ่มแก้ไขข้อมูลพนักงาน
    เข้าสู่หน้าแก้ไขข้อมูลของพนักงานต้องเจอข้อมูลพนักกงาน    0    shuhari    นารีนารถ เนรัญชร    Nareenart Narunchon    nareenart@scrum123.com    0    0    0    25000    0    750    5    salary    1500
    ใส่เงินเดือน    0    30000
    กดยืนยันการแก้ไขข้อมูล    0
    ปิด Browser

***Keywords***
เปิด Browser 
    Open Browser    ${URL_PAYMENTS}    chrome

ใส่เดือนและปีที่ต้องการดูสรุปผล
    Input Text    id=date_summary    12\t2018

กดปุ่มยืนยันดูสรุปผล
    Click Element    id=button_show_summary

เข้าสู่หน้าสรุปผลค่าจ้างเดือนและปีนั้น
    Element Text Should Be    id=title_payments    PAYMENTS

ต้องเจอพนักงานในตารางแรก
    [Arguments]    ${id}    ${name}    ${coaching}    ${training}    ${other}    ${total_incomes}    ${salary}    ${income_tax_1}    ${social_security}    ${net_salary}    ${wage}    ${income_tax_53_percentage}    ${income_tax_53}    ${net_wage}    ${net_transfer}    ${status_checking_transfer}
    Wait Until Page Contains Element    row_summary_id_${id}
    Element Text Should Be    id=member_name_th_id_${id}    ${name}
    Element Text Should Be    id=coaching_id_${id}    ${coaching}
    Element Text Should Be    id=training_id_${id}    ${training}
    Element Text Should Be    id=other_id_${id}    ${other}
    Element Text Should Be    id=total_incomes_id_${id}    ${total_incomes}
    Element Text Should Be    id=salary_id_${id}    ${salary}
    Element Text Should Be    id=income_tax_${id}_id_${id}    ${income_tax_1}
    Element Text Should Be    id=social_security_id_${id}    ${social_security}
    Element Text Should Be    id=net_salary_id_${id}    ${net_salary}
    Element Text Should Be    id=wage_id_${id}    ${wage}
    Element Text Should Be    id=income_tax_53_percentage_id_${id}    ${income_tax_53_percentage}
    Element Text Should Be    id=income_tax_53_id_${id}    ${income_tax_53}
    Element Text Should Be    id=net_wage_id_${id}    ${net_wage}
    Element Text Should Be    id=net_transfer_id_${id}    ${net_transfer}
    Select From List By Value    id=status_checking_transfer_${id}    ${status_checking_transfer}

เปลี่ยนสถานะการตรวจสอบ
    [Arguments]    ${id}    ${status}
    Select From List By Label    id=status_checking_transfer_${id}    ${status}

ใส่วันที่โอนเงิน
    [Arguments]    ${id}    ${date}
    Input Text    id=date_transfer_${id}    ${date}

ใส่หมายเหตุ
    [Arguments]    ${id}    ${comment}
    Input Text    id=comment_${id}    ${comment}

กดปุ่มเปลี่ยนสถานะ
    [Arguments]    ${id}
    Click Element    id=button_change_status_checking_transfer_id_${id}

ตรวจสอบสถานะการโอน
    [Arguments]    ${id}    ${status}    ${date}    ${comment}
    Wait Until Page Contains Element    row_summary_id_${id}
    Select From List By Value    id=status_checking_transfer_${id}    ${status}
    Textfield Should Contain    id=date_transfer_${id}    ${date}
    Textfield Should Contain    id=comment_${id}     ${comment} 

ใส่เดือนและปีที่ต้องการดูสรุปผลรายบุคคล
    [Arguments]    ${date}
    Input Text    id=date    ${date}

ใส่ชื่อที่ต้องการดูสรุปผลรายบุคคล
    [Arguments]    ${id}
    Select From List By Label    id=id    ${id}

กดปุ่มยืนยันดูสรุปค่าจ้างรายบุคคล
    Click Element    id=button_show_summary_by_id
    
เข้าสู่หน้าสรุปผลค่าจ้างรายบุคคลของเดือนและปีนั้น
    [Arguments]    ${name}    ${email}    ${overtime_rate}    ${rate_per_day}    ${rate_per_hour}    ${month}    ${full_month}    ${year}
    Element Text Should Be    id=member_name_eng    ${name}
    Element Text Should Be    id=email    ${email}
    Element Text Should Be    id=overtime_rate    ${overtime_rate}
    Element Text Should Be    id=rate_per_day    ${rate_per_day}
    Element Text Should Be    id=rate_per_hour    ${rate_per_hour}
    Element Text Should Be    id=month    ${month}
    Element Text Should Be    id=full_month    ${full_month}
    Element Text Should Be    id=year    ${year}

ชั่วโมงการทำงานทั้งหมด
    [Arguments]    ${total_hours}
    Element Text Should Be    id=thours    ${total_hours}

ค่าจ้างทั้งหมดเป็น
    [Arguments]    ${payment_wage}    ${total_coaching_customer_charging}    ${total_coaching_payment_rate}    ${total_trainig_wage}    ${total_other_wage}
    Element Text Should Be    id=payment_wage    ${payment_wage}
    Element Text Should Be    id=total_coaching_customer_charging    ${total_coaching_customer_charging}
    Element Text Should Be    id=total_coaching_payment_rate    ${total_coaching_payment_rate}
    Element Text Should Be    id=total_trainig_wage    ${total_trainig_wage}
    Element Text Should Be    id=total_other_wage    ${total_other_wage}

ใส่วันที่ต้องการเพิ่มค่าจ้างรายวัน
    [Arguments]    ${day}
    Select From List By Label    id=day    ${day}

ใส่เวลาเริ่มงานช่วงเช้า
    [Arguments]    ${time}
    Input Text    id=start_time_am    ${time}
    
ใส่เวลาจบงานช่วงเช้า
    [Arguments]    ${time}
    Input Text    id=end_time_am    ${time}

ใส่เวลาเริ่มงานช่วงบ่าย
    [Arguments]    ${time}
    Input Text    id=start_time_pm    ${time}

ใส่เวลาจบงานช่วงบ่าย
    [Arguments]    ${time}
    Input Text    id=end_time_pm    ${time}

ใส่ชั่วโมงการทำงานล่วงเวลา
    [Arguments]    ${hour}
    Input Text    id=overtime    ${hour}

ใส่ชั่วโมงการทำงานรวมของวัน
    [Arguments]    ${time}
    Input Text    id=total_hours    ${time}

ใส่ค่า Coaching Customer Charging (THB)
    [Arguments]    ${amount}
    Select From List By Label    id=coaching_customer_charging    ${amount}

ใส่ค่า Coaching Payment Rate (THB)
    [Arguments]    ${amount}
    Select From List By Label    id=coaching_payment_rate    ${amount}

ใส่ค่า Training Wage (THB)
    [Arguments]    ${amount}
    Select From List By Label    id=training_wage    ${amount}

ใส่ค่า Other Wage (THB)
    [Arguments]    ${amount}
    Select From List By Label    id=other_wage    ${amount}

ใส่บริษัทที่เป็นผู้รับผิดชอบ
    [Arguments]    ${company}
    Select From List By Label    id=company    ${company}

ใส่คำอธิบายสถานที่หรือหมายเหตุ
    [Arguments]    ${description}
    Input Text    id=description    ${description}

กดปุ่มยืนยันเพิ่มค่าจ้างรายวัน
    Click Element    id=button_add_income_item

กดปุ่มคำนวณสรุปผลค่าจ้าง
    Click Element    id=button_calculate_payment

กดปุ่มแก้ไขข้อมูลพนักงาน
    Click Element    id=button_edit_member

เข้าสู่หน้าแก้ไขข้อมูลของพนักงานต้องเจอข้อมูลพนักกงาน
    [Arguments]    ${id}    ${company}    ${name_th}    ${name_eng}    ${email}    ${overtime_rate}    ${rate_per_day}    ${rate_per_hour}    ${salary}    ${income_tax_1}    ${social_security}    ${income_tax_53_percentage}    ${status}    ${travel_expense}
    Element Text Should Be    id=company_id_${id}    ${company}
    Textfield Should Contain    id=member_name_th_id_${id}    ${name_th}
    Textfield Should Contain    id=member_name_eng_id_${id}    ${name_eng}
    Textfield Should Contain    id=email_id_${id}    ${email}
    Textfield Should Contain    id=overtime_rate_id_${id}    ${overtime_rate}
    Textfield Should Contain    id=rate_per_day_id_${id}    ${rate_per_day}
    Textfield Should Contain    id=rate_per_hour_id_${id}    ${rate_per_hour}
    Textfield Should Contain    id=salary_id_${id}    ${salary}
    Textfield Should Contain    id=income_tax_1_id_${id}    ${income_tax_1}
    Textfield Should Contain    id=social_security_id_${id}    ${social_security}
    Textfield Should Contain    id=income_tax_53_percentage_id_${id}    ${income_tax_53_percentage}
    Select From List By Value   id=status_id_${id}    ${status}
    Textfield Should Contain    id=travel_expense_id_${id}    ${travel_expense}

ใส่เงินเดือน
    [Arguments]    ${id}    ${salary}
    Input Text    id=salary_id_${id}    ${salary}

กดยืนยันการแก้ไขข้อมูล
    [Arguments]    ${id}
    Click Element    button_edit_member_id_${id}
 
ปิด Browser
    Close Browser