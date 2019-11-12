*** Settings ***
Library    SeleniumLibrary

***Variables***
${URL_PAYMENTS}    http://localhost:8080/home

*** Test Case ***
ดูผลสรุปในหน้า PAYMENTS และเปลี่ยนสถานะการโอนเงินเป็น ถูกต้อง สำเร็จ
    เปิด Browser
    กดปุ่มเข้าสู่ระบบ
    ใส่ username    logintest535@gmail.com
    ใส่ password    logintest535gmailcom
    ใส่เดือนและปีที่ต้องการดูสรุปผล    12\t2018
    กดปุ่มยืนยันดูสรุปผล    
    เข้าสู่หน้าสรุปผลค่าจ้างเดือนและปีนั้น    12-DECEMBER2018-TIMESHEET
    ต้องเจอพนักงานในตารางแรก ไม่มีสิทธิแก้ไข    1    ประธาน ด่านสกุลเจริญกิจ    75000.00    30000.00    40000.00    145000.00    80000.00    5000.00    0.00    75000.00    65000.00    10    6500.00    58500.00    133500.00    รอการตรวจสอบ
    ต้องเจอพนักงานในตารางแรก มีสิทธิแก้ไข    5    ทดสอบ เข้าสู่ระบบ    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00    5    0.00    0.00    0.00    รอการตรวจสอบ
    เปลี่ยนสถานะการตรวจสอบ    5    ถูกต้อง
    ใส่วันที่โอนเงิน    5    28/12/2018
    ใส่หมายเหตุ    5     ค่าตั๋วที่ออกไปก่อน = 169,380.00 บาท
    กดปุ่มเปลี่ยนสถานะ    5
    ใส่เดือนและปีที่ต้องการดูสรุปผล    12\t2018
    กดปุ่มยืนยันดูสรุปผล 
    ตรวจสอบสถานะการโอน    5    ถูกต้อง    28/12/2018     ค่าตั๋วที่ออกไปก่อน = 169,380.00 บาท

ดูสรุปค่าจ้างรายบุคคล และสามารถเพิ่มค่าจ้างรายวัน สำเร็จ
    ใส่เดือนและปีที่ต้องการดูสรุปผลรายบุคคล    12\t2018
    ใส่ชื่อที่ต้องการดูสรุปผลรายบุคคล     TEST LOGIN
    กดปุ่มยืนยันดูสรุปค่าจ้างรายบุคคล
    เข้าสู่หน้าสรุปผลค่าจ้างรายบุคคลเดือนและปีนั้น    12-DECEMBER2018-TIMESHEET
    เข้าสู่หน้าสรุปผลค่าจ้างรายบุคคลของเดือนและปีนั้น    Test Login    logintest535@gmail.com
    ชั่วโมงการทำงานทั้งหมด    08:00:00
    ค่าจ้างทั้งหมดเป็น    0.00    0.00    0.00    0.00    0.00
    ใส่วันที่ต้องการเพิ่มค่าจ้างรายวัน    28122018
    ใส่เวลาเริ่มงานช่วงเช้า    090000
    ใส่เวลาจบงานช่วงเช้า    120000
    ใส่เวลาเริ่มงานช่วงบ่าย    130000
    ใส่เวลาจบงานช่วงบ่าย    180000
    ใส่ค่า Coaching Customer Charging (THB)    ฿ 15,000.00
    ใส่ค่า Coaching Payment Rate (THB)    ฿ 10,000.00
    ใส่ค่า Training Wage (THB)    ฿ 0.00
    ใส่ค่า Other Wage (THB)    ฿ 0.00
    ใส่บริษัทที่เป็นผู้รับผิดชอบ    Shuhari
    ใส่คำอธิบายสถานที่หรือหมายเหตุ    Siam Chamnankit and SHR operation
    กดปุ่มยืนยันเพิ่มค่าจ้างรายวัน

แก้ไขข้อมูลพนักงาน สำเร็จ
    กดปุ่มแก้ไขข้อมูลพนักงาน
    เข้าสู่หน้าแก้ไขข้อมูลของพนักงานต้องเจอข้อมูลพนักกงาน สามาถแก้ไขได้    0    siam_chamnankit    ทดสอบ เข้าสู่ระบบ    Test Login    logintest535@gmail.com    0.00    0.00    0.00    0.00    0.00    0.00    5    salary    0.00
    ใส่เงินเดือน    0    30000
    กดยืนยันการแก้ไขข้อมูล    0
    กดปุ้มย้อนกลับ
    กดปุ่ม logout
    ปิด Browser

***Keywords***
เปิด Browser 
    Open Browser    ${URL_PAYMENTS}    chrome
    Maximize Browser Window

กดปุ่มเข้าสู่ระบบ
    Click Element    id=button_login

ใส่ username
    [Arguments]    ${username}
    Wait Until Page Contains    ลงชื่อเข้าใช้งาน
    Input Text    id=identifierId    ${username}
    Click Element    id=identifierNext

ใส่ password
    [Arguments]    ${password}
    Wait Until Page Contains    ยินดีต้อนรับ
    sleep  2s
    Input Password    name=password    ${password}
    Click Element    id=passwordNext

ใส่เดือนและปีที่ต้องการดูสรุปผล
    [Arguments]    ${date}
    Wait Until Page Contains    TIMESHEET
    Input Text    id=date_summary    ${date}

กดปุ่มยืนยันดูสรุปผล
    Click Element    id=button_show_summary

เข้าสู่หน้าสรุปผลค่าจ้างเดือนและปีนั้น
    [Arguments]    ${title}
    Element Text Should Be    id=title_timesheet    ${title}

ต้องเจอพนักงานในตารางแรก ไม่มีสิทธิแก้ไข
    [Arguments]    ${id}    ${name}    ${coaching}    ${training}    ${other}    ${total_incomes}    ${salary}    ${income_tax_1}    ${social_security}    ${net_salary}    ${wage}    ${income_tax_53_percentage}    ${income_tax_53}    ${net_wage}    ${net_transfer}    ${status_checking_transfer}
    Wait Until Page Contains Element    row_summary_id_${id}
    Element Text Should Be    id=member_name_th_id_${id}    ${name}
    Element Text Should Be    id=coaching_id_${id}    ${coaching}
    Element Text Should Be    id=training_id_${id}    ${training}
    Element Text Should Be    id=other_id_${id}    ${other}
    Element Text Should Be    id=total_incomes_id_${id}    ${total_incomes}
    Element Text Should Be    id=salary_id_${id}    ${salary}
    Element Text Should Be    id=income_tax_1_id_${id}    ${income_tax_1}
    Element Text Should Be    id=social_security_id_${id}    ${social_security}
    Element Text Should Be    id=net_salary_id_${id}    ${net_salary}
    Element Text Should Be    id=wage_id_${id}    ${wage}
    Element Text Should Be    id=income_tax_53_percentage_id_${id}    ${income_tax_53_percentage}
    Element Text Should Be    id=income_tax_53_id_${id}    ${income_tax_53}
    Element Text Should Be    id=net_wage_id_${id}    ${net_wage}
    Element Text Should Be    id=net_transfer_id_${id}    ${net_transfer}
    Element Text Should Be    id=status_checking_transfer_${id}    ${status_checking_transfer}

ต้องเจอพนักงานในตารางแรก มีสิทธิแก้ไข
    [Arguments]    ${id}    ${name}    ${coaching}    ${training}    ${other}    ${total_incomes}    ${salary}    ${income_tax_1}    ${social_security}    ${net_salary}    ${wage}    ${income_tax_53_percentage}    ${income_tax_53}    ${net_wage}    ${net_transfer}    ${status_checking_transfer}
    Wait Until Page Contains Element    row_summary_id_${id}
    Element Text Should Be    id=member_name_th_id_${id}    ${name}
    Element Text Should Be    id=coaching_id_${id}    ${coaching}
    Element Text Should Be    id=training_id_${id}    ${training}
    Element Text Should Be    id=other_id_${id}    ${other}
    Element Text Should Be    id=total_incomes_id_${id}    ${total_incomes}
    Element Text Should Be    id=salary_id_${id}    ${salary}
    Element Text Should Be    id=income_tax_1_id_${id}    ${income_tax_1}
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
    Click Button    id=button_change_status_checking_transfer_id_${id}

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
    [Arguments]    ${name}    ${email}
    Element Text Should Be    id=member_name_eng    ${name}
    Element Text Should Be    id=email    ${email}

เข้าสู่หน้าสรุปผลค่าจ้างรายบุคคลเดือนและปีนั้น
    [Arguments]    ${title}
    Element Text Should Be    id=title_timesheet_by_id    ${title}

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
    Input Text    id=day    ${day}

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
    Wait Until Element Is Visible   id=button_calculate_payment
    Click Button    id=button_calculate_payment

กดปุ่มแก้ไขข้อมูลพนักงาน
    sleep  2s
    Click Element    id=button_edit_member

เข้าสู่หน้าแก้ไขข้อมูลของพนักงานต้องเจอข้อมูลพนักกงาน สามาถแก้ไขได้
    [Arguments]    ${id}    ${company}    ${name_th}    ${name_eng}    ${email}    ${overtime_rate}    ${rate_per_day}    ${rate_per_hour}    ${salary}    ${income_tax_1}    ${social_security}    ${income_tax_53_percentage}    ${status}    ${travel_expense}
    Element Text Should Be    id=company_id_${id}    ${company}
    Element Text Should Be    id=member_name_th_id_${id}    ${name_th}
    Element Text Should Be    id=member_name_eng_id_${id}    ${name_eng}
    Element Text Should Be    id=email_id_${id}    ${email}
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
    Scroll Element Into View    id=button_edit_member_id_${id}
    Click Element    id=button_edit_member_id_${id}

กดปุ้มย้อนกลับ
    Click Element    id=button_to_top
    Scroll Element Into View    id=button_back
    Wait Until Element Is Visible   id=button_back
    Click Element    id=button_back

กดปุ่ม logout
    Click Element    id=button_logout

ปิด Browser
    Close Browser