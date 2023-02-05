import React, { useState, useEffect } from 'react';
import Paper from '@mui/material/Paper';
import CustomerShow from './CustomerShow';
import CustomerEdit from './CustomerEdit';
import { CustomerInterface } from '../../../interfaces/CustomerUI';


const steps = [
    'Profile Information',
    'Create Account',
    'Sign - up Complete',
];

const defaultCreate = {
    ID: 0,
    Name: "",
    CAREER_ID: 0,
    Phone: "",
};

function CreateForm2 (){

    const [formCreate, setFormCreate] = useState(defaultCreate);
    const [Customer, setCustomer] = useState<Partial<CustomerInterface>>({});
    const [activeStep, setActiveStep] = React.useState(0);


    const PageDisplay = () => {
        if (activeStep === 0) {
            return <CustomerShow activeStep={activeStep} setActiveStep={setActiveStep} formCreate={formCreate} setFormCreate={setFormCreate}  />

        } else if (activeStep === 1) {
            return <CustomerEdit formCreate={formCreate} setFormCreate={setFormCreate} activeStep={activeStep} setActiveStep={setActiveStep} steps={steps} />

        } 
        //   else if (activeStep === 2) {
        //     return <EditDataReview Cusomer={Cusomer} setActiveStep={setActiveStep} />
        // }
    }


    return (
        <Paper
            sx={{
                backgroundColor: "#182e3e",
                height: '1500px'
            }}>
            <form className='form-container'>
                <div className='text-start'>{PageDisplay()}</div>
            </form >
        </Paper>
    );
}
export default CreateForm2
