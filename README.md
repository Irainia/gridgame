# GridGame <br><br>
# Problem Description
This is Grid Shape game. At the start, the player is presented with board sized 10x10. Then, the computer prepare <i>several</i> shapes unknown to the player. The number, order, type of shape, and the size of shape are unknown to the player. The player should put <b>all</b> the shapes to the board. To put each shape, the player specifies its <i>coordinate</i>.

<b>RESTRICTION</b><br>
- If even one of the shape is overlapped with the other shape(s) or is out of the board, then the player lose.
- If the player put all the shapes to the board successfully, then the board will expand either to buttom or right of the board randomly. The process is repeated with the doubled size of the board and the previously filled board is kept and used.
- Board is sized <i>ROW x COLUMN</i> with <i>ROW</i> is the number of row top-down increasing and <i>COLUMN</i> is the number of column left-right increasing.
- The score is calculated the number of unit filled on the board and the highest score is saved to be loaded later.