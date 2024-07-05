/*
 * Time Tracker
 * API for control user's working time
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.ModelsGetAllUsersResponse;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for GetInfoApi
 */
@Ignore
public class GetInfoApiTest {

    private final GetInfoApi api = new GetInfoApi();

    
    /**
     * GetUserInfo from API
     *
     * get user info
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void usersGetTest() throws ApiException {
        ModelsGetAllUsersResponse response = api.usersGet();

        // TODO: test validations
    }
    
}