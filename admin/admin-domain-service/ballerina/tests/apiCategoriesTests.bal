//
// Copyright (c) 2023, WSO2 LLC. (http://www.wso2.com).
//
// WSO2 LLC. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

import ballerina/test;

APICategory  apiCategory = {name: "MyCat1", description: "My Desc 1", id: "01ed9241-2d5d-1b98-8ecb-40f85676b081", numberOfAPIs: 2};

@test:Config {}
function addAPICategoryTest() {
    APICategory payload = {name: "MyCat1", description: "My Desc 1"};
    CreatedAPICategory|APKError createdApiCategory = addAPICategory(payload);
    if createdApiCategory is CreatedAPICategory {
        test:assertTrue(true,"API Category added successfully");
        apiCategory.id = createdApiCategory.body.id;
    } else if createdApiCategory is APKError {
        test:assertFail("Error occured while adding API Category");
    }
}

@test:Config  {dependsOn: [addAPICategoryTest]}
function addAPICategoryTestNegative1() {
    APICategory payload = {name: "MyCat1", description: "My Desc 1"};
    //API Category Name alrady exisitng
    CreatedAPICategory|APKError createdApiCategory = addAPICategory(payload);
    if createdApiCategory is CreatedAPICategory {
        test:assertFail("API Category added successfully");
    } else if createdApiCategory is APKError {
        test:assertTrue(true, "Error occured while adding API Category");
    }
}

@test:Config {dependsOn: [addAPICategoryTest]}
function getAllCategoryListTest() {
    APICategoryList|APKError apiCategoryList = getAllCategoryList();
    if apiCategoryList is APICategoryList {
        test:assertTrue(true,"API Category list retrieved successfully");
    } else if apiCategoryList is APKError {
        test:assertFail("Error occured while retrieving API Category List");
    }
}

@test:Config {dependsOn: [getAllCategoryListTest]}
function updateAPICategoryTest() {
    APICategory payload = {name: "MyCat1", description: "My Desc 1 new"};
    string? catId = apiCategory.id;
    if catId is string {
        APICategory|NotFoundError|APKError createdApiCategory = updateAPICategory(catId,payload);
        if createdApiCategory is APICategory {
            test:assertTrue(true,"API Category updated successfully");
        } else if createdApiCategory is NotFoundError {
            test:assertFail("Not Found Error");
        } else if createdApiCategory is APKError {
            test:assertFail("Error occured while adding API Category");
        }
    } else {
        test:assertFail("Category ID isn't a string");
    }
}

@test:Config {dependsOn: [updateAPICategoryTest]}
function updateAPICategoryTestNegative1() {
    // Exisiting API Category is not found
    APICategory payload = {name: "MyCat1", description: "My Desc 1 new"};
    string? catId = apiCategory.id;
    if catId is string {
        APICategory|NotFoundError|error createdApiCategory = updateAPICategory("01ed9241-2d5d-1b98-8ecb-40f85676b081",payload);
        if createdApiCategory is APICategory {
            test:assertFail("API Category updated successfully");
        } else if createdApiCategory is NotFoundError {
            test:assertTrue(true, "Not Found Error");
        } else if createdApiCategory is APKError {
            test:assertFail("Error occured while adding API Category");
        }
    } else {
        test:assertFail("Category ID isn't a string");
    }
}

@test:Config {dependsOn: [updateAPICategoryTestNegative1]}
function updateAPICategoryTestNegative2() {
    // Another API Category by same name exists
    APICategory payloadOther = {name: "MyCat2", description: "My Desc 1"};
    CreatedAPICategory|APKError anotherApiCategory = addAPICategory(payloadOther);

    // New Name
    APICategory payload = {name: "MyCat2", description: "My Desc 1 new"};
    string? catId = apiCategory.id;
    if catId is string {
        APICategory|NotFoundError|APKError createdApiCategory = updateAPICategory(catId,payload);
        if createdApiCategory is APICategory {
            test:assertFail("API Category updated successfully");
        } else if createdApiCategory is NotFoundError {
            test:assertFail("Not Found Error");
        } else if createdApiCategory is APKError {
            test:assertTrue(true,"Error occured while adding API Category");
        }
    } else {
        test:assertFail("Category ID isn't a string");
    }
}

@test:Config {dependsOn: [updateAPICategoryTestNegative2]}
function removeAPICategoryTest(){
    string? catId = apiCategory.id;
    if catId is string {
        APKError|string status = removeAPICategory(catId);
        if status is string {
        test:assertTrue(true, "Successfully deleted API Category");
        } else if status is  APKError {
            test:assertFail("Error occured while deleting API Category");
        }
    } else {
        test:assertFail("Category ID isn't a string");
    }
}
