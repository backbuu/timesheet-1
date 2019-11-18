*** Settings ***
Library    SeleniumLibrary
Suite Setup    Open Browser    ${LOCAL_URL}    ${BROWSER} 
Test Setup    Go To    ${LOCAL_URL}
Suite Teardown    Close Browser

Resource    resource_success.robot

***Variables***
${URL}    http://localhost:8080/home

*** Test Case ***
แก้ไขข้อมูลพนักงาน สำเร็จ
    เข้าสู่ระบบ 
    ใส่เดือนและปีที่ต้องการดูสรุปผล    12\t2018    12-DECEMBER2018-TIMESHEET
    เข้าสู่หน้าแก้ไขข้อมูลพนักงาน    employeeID001
    แสดงหน้าแก้ไขข้อมูลพนักงาน    0    Siam Chamnankit    ประธาน ด่านสกุลเจริญกิจ    Prathan Dansakulcharoenkit    prathan@scrum123.com    15000.00    1875    80000.00    5000.00    0.00    10    wage    0.00
    แก้ไขข้อมูลเงินเดือนพนักงาน    0    100000.00
    แสดงหน้าแก้ไขข้อมูลพนักงาน    0    Siam Chamnankit    ประธาน ด่านสกุลเจริญกิจ    Prathan Dansakulcharoenkit    prathan@scrum123.com    15000.00    1875    100000.00    5000.00    0.00    10    wage    0.00