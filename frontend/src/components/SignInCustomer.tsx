import React, { useState } from "react";
import Box from "@mui/material/Box";
import "./Login.css"
import { Button, Snackbar, TextField, Typography } from "@mui/material";

import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { SigninInterface } from "../interfaces/ISignin";
import { Login } from "../services/HttpClientService";

// import { Login } from "../services/HttpClientService";
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

const theme = createTheme();

const defaultSignin = {
    Email: "",
    Password: "",
};

function SignInCustomer() {
    const [signin, setSignin] = useState(defaultSignin)
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    console.log(signin);


    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof signin;
        const { value } = event.target;
        setSignin({ ...signin, [id]: value });
    };

    const handleClose = (
        event?: React.SyntheticEvent | Event,
        reason?: string
    ) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    const submit = async () => {
        let res = await Login(signin);
        if (res) {
            setSuccess(true);
            setTimeout(() => {
                window.location.reload();
            }, 1000);
        } else {
            setError(true);
        }
    };
    return (

        <div >
            <Snackbar
                open={success}
                autoHideDuration={3000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="success">
                    เข้าสู่ระบบสำเร็จ
                </Alert>
            </Snackbar>
            <Snackbar
                open={error}
                autoHideDuration={3000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="error">
                    อีเมลหรือรหัสผ่านไม่ถูกต้อง
                </Alert>
            </Snackbar>
            <Box className="loginform" id="loginform">
                <Typography component="h1" variant="h5">
                    Sign In For Customer
                </Typography>
                <div id="textField">
                    <TextField
                        variant="standard"
                        required
                        fullWidth
                        id="Email"
                        label="Email Address"
                        name="email"
                        autoComplete="email"
                        autoFocus
                        value={signin.Email}
                        onChange={(event) =>
                            setSignin(({ ...signin, Email: event.target.value }))}
                    />
                </div>

                <div id="textField">
                    <TextField
                        className="textField"
                        variant="standard"
                        required
                        fullWidth
                        name="password"
                        label="Password"
                        type="password"
                        id="Password"
                        autoComplete="current-password"
                        value={signin.Password}
                        onChange={(event) =>
                            setSignin(({ ...signin, Password: event.target.value }))}
                    />
                </div>


                <Button
                    id="buttonSignInCustomert"
                    // className="buttonLogin"
                    type="submit"
                    variant="contained"
                    sx={{ mt: 3, mb: 2 }}
                    onClick={submit}
                >
                    Sign In
                </Button>
                <Typography className="alternativeLogin">
                    Or Sign In
                </Typography>
                <Button
                    id="buttonGoSignInTeshPage"
                    type="submit"
                    variant="contained"
                    sx={{ mt: 3, mb: 2 }}
                // onClick={submit}
                >
                    หน้า sing in สำหรับช่าง
                </Button>
                <Typography className="alternativeSignUp">
                    Don't have an account yet?
                    <Button
                    // onClick={submit}
                    >
                        Sign Upppppp
                    </Button>
                </Typography>

            </Box>
        </div>
    );
}

export default SignInCustomer;