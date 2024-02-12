/**
 * This file was auto-generated by Fern from our API Definition.
 */
import * as serializers from "../../..";
import * as SeedTrace from "../../../../api";
import * as core from "../../../../core";
export declare const TestSubmissionStatus: core.serialization.Schema<serializers.TestSubmissionStatus.Raw, SeedTrace.TestSubmissionStatus>;
export declare namespace TestSubmissionStatus {
    type Raw = TestSubmissionStatus.Stopped | TestSubmissionStatus.Errored | TestSubmissionStatus.Running | TestSubmissionStatus.TestCaseIdToState;
    interface Stopped {
        type: "stopped";
    }
    interface Errored {
        type: "errored";
        value: serializers.ErrorInfo.Raw;
    }
    interface Running {
        type: "running";
        value: serializers.RunningSubmissionState.Raw;
    }
    interface TestCaseIdToState {
        type: "testCaseIdToState";
        value: Record<string, serializers.SubmissionStatusForTestCase.Raw>;
    }
}