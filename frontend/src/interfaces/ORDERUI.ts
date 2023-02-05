import { AddressInterface } from "./AddressUI";
import { DeviceInterface } from "./IDevice";

export interface ORDERInterface {
    ID: number,
    Date_time: Date | null;
    Reason: string;
    Limit: number;
    CASEID: number;
    CASE: CASEInterface
    DeviceID: number;
    Device: DeviceInterface
    AddressID: number;
    Address: AddressInterface
    StateID :number;
    State: StateInterface
}

export interface CASEInterface {
    ID: number,
    Case_text: string,
    Level_case:string,
}

export interface StateInterface {
    ID: number,
    State: string,
}


