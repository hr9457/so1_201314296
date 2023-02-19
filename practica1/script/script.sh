#!/bin/bash

content=$(curl -L http://192.168.0.10:4000/historial)

operaciones=$(echo $content | jq '.. | .operador? //empty')
sumas=$(echo $operaciones | grep -o "+" | wc -l )
restas=$(echo $operaciones | grep -o "-" | wc -l )
multi=$(echo $operaciones | grep -o "*" | wc -l )
div=$(echo $operaciones | grep -o "/" | wc -l )
err=$(echo $operaciones | grep -o "Error" | wc -l )
total=$(($sumas+$restas+$multi+$div+$err))

echo "********* LOGS REGISTRADOS *********************"
echo "${total}"

echo "************* ERRORES **********************"
echo "${err}"

echo "************* OPERACIONES **********************"
echo "Operacion                 Totales"
echo "Total sumas:              ${sumas}"
echo "Total restas:             ${restas}"
echo "Total multiplicaciones:   ${multi}"
echo "Total divisiones:         ${div}"


echo "********* HISTORIAL OPERACIONES ******************"
echo "${content}"
