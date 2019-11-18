*** Settings ***
Library    SeleniumLibrary
Suite Setup    Open Browser    ${LOCAL_URL}    ${BROWSER} 
Test Setup    Go To    ${LOCAL_URL}
Suite Teardown    Close Browser

Resource    resource_success.robot

***Variables***
${URL}    http://localhost:8080/home

*** Test Case ***
เข้าสู่เว็บไซต์
    เข้าสู่ระบบ 
    ใส่เดือนและปีที่ต้องการดูสรุปผล    12\t2018    12-DECEMBER2018-TIMESHEET

แก้ไขข้อมูลพนักงาน Prathan Dansakulcharoenkit เปลี่ยนเงินเดือนเป็น 100,000.00 บาทสำเร็จ
    เข้าสู่หน้าแก้ไขข้อมูลพนักงาน    employeeID001
    แสดงหน้าแก้ไขข้อมูลพนักงาน    0    Siam Chamnankit    ประธาน ด่านสกุลเจริญกิจ    Prathan Dansakulcharoenkit    prathan@scrum123.com    15000.00    1875    80000.00    5000.00    0.00    10    wage    0.00
    แสดงหน้าแก้ไขข้อมูลพนักงาน    1    SHU HA RI    ประธาน ด่านสกุลเจริญกิจ    Prathan Dansakulcharoenkit    prathan@scrum123.com    15000.00    1875    0.00    0.00    0.00    10    wage    0.00
    แก้ไขข้อมูลพนักงาน    0    15000.00    1875    100000.00    5000.00    0.00    10    ค่าจ้างรายวัน (wage)    0.00
    แสดงหน้าแก้ไขข้อมูลพนักงาน    0    Siam Chamnankit    ประธาน ด่านสกุลเจริญกิจ    Prathan Dansakulcharoenkit    prathan@scrum123.com    15000.00    1875    100000.00    5000.00    0.00    10    wage    0.00

แก้ไขข้อมูลพนักงาน Nareenart Narunchon เปลี่ยนเงินเดือนเป็น 30,000.00 บาทสำเร็จ
    เข้าสู่หน้าแก้ไขข้อมูลพนักงาน    employeeID002
    แสดงหน้าแก้ไขข้อมูลพนักงาน    0    SHU HA RI    นารีนารถ เนรัญชร    Nareenart Narunchon    nareenart@scrum123.com    0.00    0    25000.00    0.00    0.00    5    salary    1500.00
    แก้ไขข้อมูลพนักงาน    0    0.00    0    30000.00    0.00    0.00    5    เงินเดือน (salary)    1500.00
    แสดงหน้าแก้ไขข้อมูลพนักงาน    0    SHU HA RI    นารีนารถ เนรัญชร    Nareenart Narunchon    nareenart@scrum123.com    0.00    0    30000.00    0.00    0.00    5    salary    1500.00