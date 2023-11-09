// Import the necessary k6 libraries
import http from 'k6/http';
import { check } from 'k6';
import file from 'k6/x/file';

// Define file paths for failed and successful responses
const failedResponseFilePath = './files/failed.csv'
const successResponseFilePath = './files/success.csv'

// Base URL for the API
const baseUrl = 'https://fakestoreapi.com/products';

// Define the options for the test
export let options = {
	iterations: 5,
	vus: 1,
    duration: '6s',
};

// Setup function to prepare for the test
export function setup() {
    // Clear all the previous data
    file.clearFile(successResponseFilePath)
    file.clearFile(failedResponseFilePath)

    // Write the first line to the file
    file.writeString(successResponseFilePath, 'New Run.\n');
    file.writeString(failedResponseFilePath, 'New Run.\n');
}

// Default function that is called in each iteration of the test
export default function() {
    // Get the current iteration number
    let i = __ITER;

    // Define the headers for the request
    let headers = {
            'Content-Type': 'application/json',
        }

    // Send a GET request to the API
    let getProduct = http.get(`${baseUrl}/${i}`, {
            headers: headers
        });

    // Check the response for certain conditions
    check(getProduct, {
        'status is 200': (r) => r.status === 200,
        'response time < 200ms': (r) => r.timings.duration < 200
    });

    // If the status is not 200, log the failure and append it to the file
    if (getProduct.status !== 200) {
        file.appendString(failedResponseFilePath,`Failed API Response,${getProduct.status},${getProduct.timings.duration}\n`)
    }
    // If the status is 200, log the success and append it to the file
    else if(getProduct.status === 200){
        file.appendString(successResponseFilePath, `Successful Response,${getProduct.status},${getProduct.timings.duration}\n`)
    }
}

// Teardown function to clean up after the test
export function teardown() {
    // Additional cleanup or actions can be added here
    console.log("Teardown function executed.");
}
