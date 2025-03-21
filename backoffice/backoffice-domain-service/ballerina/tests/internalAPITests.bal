//
// Copyright (c) 2022, WSO2 LLC. (http://www.wso2.com).
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

import ballerina/log;
import ballerina/test;

@test:Config {}
function createAPITest() {
    APIBody body = {
        "apiProperties":{
            "id": "01ed75e2-b30b-18c8-wwf2-25da7edd2231",
            "name":"PizzaShask",
            "context":"pizzssa",
            "version":"1.0.0",
            "provider":"admin",
            "lifeCycleStatus":"CREATED",
            "type":"HTTP"
        },
        "Definition" : {	  
        "openapi": "3.0.0",
        "info": {
            "title": "Sample API",
            "description": "Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.",
            "version": "0.1.9"
        },
        "servers": [
            {
            "url": "http://api.example.com/v1",
            "description": "Optional server description, e.g. Main (production) server"
            },
            {
            "url": "http://staging-api.example.com",
            "description": "Optional server description, e.g. Internal staging server for testing"
            }
        ],
        "paths": {
            "/users": {
            "get": {
            "summary": "Returns a list of users.",
            "description": "Optional extended description in CommonMark or HTML.",
            "responses": {
            "200": {
                "description": "A JSON array of user names",
                "content": {
                "application/json": {
                    "schema": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                    }
                }
                }
            }
            }
            }
            }
        }
        }
    };
    API|error createdAPI = createAPI(body,"carbon.super");
    if createdAPI is API {
        test:assertTrue(true, "Successfully created API");
    } else if createdAPI is  error {
        log:printError(createdAPI.toString());
        test:assertFail("Error occured while creating API");
    }

}

@test:Config {}
function updateInternalAPITest() {
    APIBody updateBody = {
            "apiProperties":{
                "name":"PizzaShask",
                "context":"pizzssa",
                "version":"1.0.0",
                "provider":"admin",
                "lifeCycleStatus":"CREATED",
                "type":"HTTP"
            },
            "Definition" : {	  
            "openapi": "3.0.0",
            "info": {
                "title": "Sample API Change",
                "description": "Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.",
                "version": "0.1.9"
            },
            "servers": [
                {
                "url": "http://api.example.com/v1",
                "description": "Optional server description, e.g. Main (production) server"
                },
                {
                "url": "http://staging-api.example.com",
                "description": "Optional server description, e.g. Internal staging server for testing"
                }
            ],
            "paths": {
                "/users": {
                "get": {
                "summary": "Returns a list of users.",
                "description": "Optional extended description in CommonMark or HTML.",
                "responses": {
                "200": {
                    "description": "A JSON array of user names",
                    "content": {
                    "application/json": {
                        "schema": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                        }
                    }
                    }
                }
                }
                }
                }
            }
            }
        };
        API|error updateAPI = updateAPI_internal("01ed75e2-b30b-18c8-wwf2-25da7edd2231", updateBody, "carbon.super");
            if updateAPI is API {
                test:assertTrue(true, "Successfully updtaing API");
            } else if updateAPI is  error {
                log:printError(updateAPI.toString());
                test:assertFail("Error occured while updtaing API");
        }
}

@test:Config {}
function updateAPIDefinitionTest() {
    APIDefinition1 apiDefinition = {
      "Definition" : {	  
            "openapi": "3.0.0",
            "info": {
                "title": "Sample API Change",
                "description": "Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.",
                "version": "0.1.9"
            },
            "servers": [
                {
                "url": "http://api.example.com/v1",
                "description": "Optional server description, e.g. Main (production) server"
                },
                {
                "url": "http://staging-api.example.com",
                "description": "Optional server description, e.g. Internal staging server for testing"
                }
            ],
            "paths": {
                "/users": {
                "get": {
                "summary": "Returns a list of users.",
                "description": "Optional extended description in CommonMark or HTML.",
                "responses": {
                "200": {
                    "description": "A JSON array of user names",
                    "content": {
                    "application/json": {
                        "schema": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                        }
                    }
                    }
                }
                }
                }
                }
            }  
        }
    };
    APIDefinition1|error? updateAPI = updateDefinition(apiDefinition, "01ed75e2-b30b-18c8-wwf2-25da7edd2231");
        if updateAPI is API {
            test:assertTrue(true, "Successfully updtaing API Definition");
        } else if updateAPI is  error {
            log:printError(updateAPI.toString());
            test:assertFail("Error occured while updtaing API Definition");
    }
}
