import React, { useCallback, useContext, useEffect, useState } from "react";
import axios from "axios";
import { AuthContext } from "../../../context/auth.context";
import {
	Paper,
	Table,
	TableBody,
	TableCell,
	TableContainer,
	TableHead,
	TableRow,
	Typography,
} from "@mui/material";
import UserBlock from "./block/user.block";

const UserServicePage = () => {
	const auth = useContext(AuthContext);
	const [users, setUsers] = useState([]);

	const GetUsers = useCallback(async () => {
		axios
			.get("http://localhost:8000/getallusers", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setUsers(res.data);
			});
	}, []);

	useEffect(() => {
		GetUsers();
	}, [GetUsers]);

	return (
		<div>
			<Typography variant="h3">Пользователи</Typography>
			<TableContainer component={Paper}>
				<Table
					sx={{ minWidth: 500 }}
					padding="none"
					aria-label="simple table"
				>
					<TableHead>
						<TableRow>
							<TableCell>Имя пользователя</TableCell>
							<TableCell align="center">почта</TableCell>
							<TableCell align="center">
								Тип пользователя
							</TableCell>
							<TableCell align="center">
								Роль пользователя
							</TableCell>
						</TableRow>
					</TableHead>
					<TableBody>
						{users.map((user) => (
							<UserBlock user={user} />
						))}
					</TableBody>
				</Table>
			</TableContainer>
		</div>
	);
};

export default UserServicePage;
