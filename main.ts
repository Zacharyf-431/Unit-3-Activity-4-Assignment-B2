/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-11-21
 * @fileoverview: This file calculates the cost of a car with selected features
 */

// Assign constants
const basePrice: number = 28000;
const floorMat: number = 800;
const navSystem: number = 1000;
const heatedSeats: number = 800;
const warrantyY: number = 2800;

// TOTAL COST
let total: number = basePrice;

// OUTPUT BASE PRICE 
console.log("Base Price".padEnd(80) + basePrice.toFixed(2).padStart(10));

// INPUT
let mats: string | null = prompt("Would you like floor mats? (yes/no): ");
let nav: string | null = prompt("Would you like a navigation system? (yes/no): ");
let seats: string | null = prompt("Would you like heated leather seats? (yes/no): ");
let warranty: string | null = prompt("Would you like a 5-year extended warranty? (yes/no): ");

if (mats === null || nav === null || seats === null || warranty === null) {
  console.log("Invalid input");
} else {

  let m = mats.toLowerCase();
  let n = nav.toLowerCase();
  let h = seats.toLowerCase();
  let w = warranty.toLowerCase();

  // PROCESS
  if (m == "yes") {
    console.log("Floor mats".padEnd(80) + floorMat.toFixed(2).padStart(10));
    total += floorMat;
  }

  if (n == "yes") {
    console.log("Navigation system".padEnd(80) + navSystem.toFixed(2).padStart(10));
    total += navSystem;
  }

  if (h == "yes") {
    console.log("Heated leather seats".padEnd(80) + heatedSeats.toFixed(2).padStart(10));
    total += heatedSeats;
  }

  if (w == "yes") {
    console.log("5-Year extended warranty".padEnd(80) + warrantyY.toFixed(2).padStart(10));
    total += warrantyY;
  }

  // TAX + FINAL
  let tax = total * 0.13;
  let finalCost = total + tax;

  console.log("13% Taxes".padEnd(80) + tax.toFixed(2).padStart(10));
  console.log("Final cost of car".padEnd(80) + finalCost.toFixed(2).padStart(10));
}

console.log("\nDone.");