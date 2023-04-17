import {
	Box,
	Grid,
	Stack,
	TableCell,
	TableRow,
	Typography,
} from "@mui/material";

import React from "react";

const ItemBlock = (props) => {
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
		</TableRow>
	);
};

export default ItemBlock;

{
	/* <Grid>
				<Typography>{props.name}</Typography>
				<br />
				<Typography>{props.id}</Typography>
			</Grid>
			<Grid>
				<Typography>
					Категория: {props.c_name}; Единица продажи: {props.dim}{" "}
				</Typography>
			</Grid>
			<Grid>
				<Typography>Цена за единицу: {props.pfu}</Typography>
			</Grid> */
}
