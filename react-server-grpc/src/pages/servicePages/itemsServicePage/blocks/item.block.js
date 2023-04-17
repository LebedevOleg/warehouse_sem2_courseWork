import {
	Box,
	Button,
	Grid,
	Stack,
	TableCell,
	TableRow,
	Typography,
} from "@mui/material";

import React from "react";

const ItemServiceBlock = (props) => {
	return (
		<TableRow key={props.product.id}>
			<TableCell align="center">
				<Stack direction="row" spacing={3}>
					<Typography variant="subtitle2">
						{props.product.id}
					</Typography>
					<Typography variant="h6">{props.product.name}</Typography>
				</Stack>
			</TableCell>
			<TableCell align="center">
				{props.product.c_name}; Единица продажи: {props.product.dim}
			</TableCell>
			<TableCell align="center">
				Цена за единицу: {props.product.pfu}
			</TableCell>
			<TableCell align="center">
				<Button>Изменить</Button>
				<Button>Удалить</Button>
			</TableCell>
		</TableRow>
	);
};

export default ItemServiceBlock;
