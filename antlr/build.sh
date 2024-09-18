#!/usr/bin/env bash

export antlr='java -jar /home/ankan/Documents/antlr-4.13.2-complete.jar'

$antlr -Dlanguage=Go /home/ankan/Documents/git/me/9tail/antlr/NinetailLexer.g4 -o parser
$antlr -Dlanguage=Go /home/ankan/Documents/git/me/9tail/antlr/NinetailParser.g4 -visitor -o parser
