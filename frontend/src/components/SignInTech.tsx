import React from "react"
import Box from "@mui/material/Box";
import "./Login.css"
import { Button, TextField, Typography } from "@mui/material";


function SignInTech() {
    return (

        <div >
            <Box className="loginform">
                <Typography component="h1" variant="h5">
                    Sign In For Tech
                </Typography>
                <div id="textField">
                    <TextField
                        variant="standard"
                        required
                        fullWidth
                        id="UserName"
                        label="User Name"
                        name="User Name"
                        autoComplete="User Name"
                        autoFocus
                    // value={signin.Email || ""}
                    // onChange={handleInputChange}
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
                    // value={signin.Password || ""}
                    // onChange={handleInputChange}
                    />
                </div>
                <Button
                    className="buttonLogin"
                    type="submit"
                    variant="contained"
                    sx={{ mt: 3, mb: 2 }}
                // onClick={submit}
                >
                    Sign In
                </Button>
                <Typography className="alternativeLogin">
                    Or Sign In 
                </Typography>
                <Button
                    className="buttonLogin"
                    type="submit"
                    variant="contained"
                    sx={{ mt: 3, mb: 2 }}
                // onClick={submit}
                >
                    หน้า sing in สำหรับลูกค้า
                </Button>
            </Box>
        </div>
    );
}

export default SignInTech;