/*
 * Oxidrive API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 *
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct InstanceSetupRequest {
    #[serde(rename = "admin")]
    pub admin: models::InstanceSetupRequestAdmin,
}

impl InstanceSetupRequest {
    pub fn new(admin: models::InstanceSetupRequestAdmin) -> InstanceSetupRequest {
        InstanceSetupRequest { admin }
    }
}
