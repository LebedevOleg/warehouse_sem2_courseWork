import { Button, Stack, TableCell, TableRow, Typography } from "@mui/material";

import React from "react";
import UserChangeModal from "../modal/userChange.modal";

const UserBlock = ({ user }) => {
	return (
		<TableRow key={user.id}>
			<TableCell align="center">
				<Stack direction="row" spacing={3}>
					<Typography variant="subtitle2">{user.id}</Typography>
					<Typography variant="h6">{user.name}</Typography>
				</Stack>
			</TableCell>
			<TableCell align="center">{user.email}</TableCell>
			<TableCell align="center">{user.type}</TableCell>
			<TableCell align="center">{user.role}</TableCell>
			<TableCell align="center">
				<UserChangeModal user={user} />
				<Button>Удалить</Button>
			</TableCell>
		</TableRow>
	);
};

export default UserBlock;
