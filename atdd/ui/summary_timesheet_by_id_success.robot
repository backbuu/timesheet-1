*** Settings ***
Library    SeleniumLibrary
Suite Setup    Open Browser    ${LOCAL_URL}    ${BROWSER} 
Test Setup    Go To    ${LOCAL_URL}
Suite Teardown    Close Browser

Resource    resource_success.robot

***Variables***
${URL}    http://localhost:8080/home

*** Test Case ***
ดูสรุปค่าจ้างรายบุคคล และสามารถเพิ่มค่าจ้างรายวัน สำเร็จ
    เข้าสู่ระบบ 
    ใส่เดือนและปีที่ต้องการดูสรุปผล    12\t2018    12-DECEMBER2018-TIMESHEET
    เลือกแสดงรายละเอียดสรุปผลรายบุคคล    employeeID001
    แสดงข้อมูลพนักงานในหน้าสรุปผลค่าจ้างรายบุคคลของเดือนและปี   12-DECEMBER2018-TIMESHEET    Prathan Dansakulcharoenkit    prathan@scrum123.com    15,000.00    1875    12    December    2018
    แสดงสรุปผลค่าจ้างรายบุคคลทั้งหมดเป็น    144:00:00    ฿ 185,000.00    ฿ 15,000.00	฿ 75,000.00	   ฿ 70,000.00    ฿ 40,000.00	
    เพิ่มค่าจ้างรายวันในส่วนของ Coaching (THB)    28122018    ฿ 15,000.00    ฿ 10,000.00    Shu Ha Ri    Siam Chamnankit and SHR operation
    เพิ่มค่าจ้างรายวันในส่วนของ Training Wage (THB)    29122018    ฿ 10,000.00    Shu Ha Ri    Siam Chamnankit and SHR operation
    เพิ่มค่าจ้างรายวันในส่วนของ Other Wage (THB)    30122018    ฿ 10,000.00    Shu Ha Ri    Siam Chamnankit and SHR operation
    แสดงสรุปผลค่าจ้างรายบุคคลทั้งหมดเป็น    168:00:00    ฿ 215,000.00    ฿ 30,000.00	฿ 85,000.00	   ฿ 80,000.00    ฿ 50,000.00	