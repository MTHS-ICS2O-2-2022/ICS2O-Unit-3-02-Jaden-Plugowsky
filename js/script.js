// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

function calculatePressed() {
  //Input through Text Fields
  const baseLength = parseFloat(document.getElementById("base-length").value)
  const baseWidth = parseFloat(document.getElementById("base-width").value)
  const height = parseFloat(document.getElementById("height").value)

  //Process
  const volume = (baseLength * baseWidth * height) / 3

  //Output
  document.getElementById("answer").innerHTML =
    "The area of the Right Rectangular Pyramid is: " + volume.toFixed(2) + "cm³"
}
