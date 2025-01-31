/**
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.apache.inlong.tubemq.server.common.utils;

import org.apache.inlong.tubemq.corebase.TErrCodeConstants;

public class ProcessResult {
    public boolean success = true;
    public int errCode = TErrCodeConstants.SUCCESS;
    public String errInfo = "";
    public Object retData1 = null;

    public ProcessResult() {

    }

    public ProcessResult(ProcessResult other) {
        this.success = other.success;
        this.errCode = other.errCode;
        this.errInfo = other.errInfo;
        this.retData1 = other.retData1;
    }

    public ProcessResult(Object retData) {
        this.success = true;
        this.retData1 = retData;
    }

    public ProcessResult(int errCode, String errInfo) {
        this.success = false;
        this.errCode = errCode;
        this.errInfo = errInfo;
    }

    public void setFailResult(int errCode, final String errMsg) {
        this.success = false;
        this.errCode = errCode;
        this.errInfo = errMsg;
        this.retData1 = null;
    }

    public void setFailResult(final String errMsg) {
        this.success = false;
        this.errCode = TErrCodeConstants.BAD_REQUEST;
        this.errInfo = errMsg;
        this.retData1 = null;
    }

    public void setSuccResult(Object retData) {
        this.success = true;
        this.errInfo = "Ok!";
        this.errCode = TErrCodeConstants.SUCCESS;
        this.retData1 = retData;
    }

    public boolean isSuccess() {
        return success;
    }

    public int getErrCode() {
        return errCode;
    }

    public String getErrInfo() {
        return errInfo;
    }

    public Object getRetData() {
        return retData1;
    }

    public void setRetData(Object retData) {
        this.retData1 = retData;
    }

    public void clear() {
        this.success = true;
        this.errCode = TErrCodeConstants.SUCCESS;
        this.errInfo = "";
        this.retData1 = null;
    }
}
