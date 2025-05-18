namespace MergeOrdenado.Cli.Test;

public class UnitTest1
{
    [Fact]
    public void MergeOrdenado_MergesTwoSortedArrays_Correctly()
    {
        // Arrange
        int[] array01 = { 1, 3, 5, 7, 9 };
        int[] array02 = { 2, 3, 6, 8, 10 };
        int[] expected = { 1, 2, 3, 3, 5, 6, 7, 8, 9, 10 };

        // Act
        int[] result = MergeOrdenado.Cli.Program.MergeOrdenado(array01, array02);

        // Assert
        Assert.Equal(expected, result);
    }

    [Fact]
    public void MergeOrdenado_ReturnsFirstArray_WhenSecondIsEmpty()
    {
        // Arrange
        int[] array01 = { 1, 2, 3 };
        int[] array02 = { };
        int[] expected = { 1, 2, 3 };

        // Act
        int[] result = MergeOrdenado.Cli.Program.MergeOrdenado(array01, array02);

        // Assert
        Assert.Equal(expected, result);
    }

    [Fact]
    public void MergeOrdenado_ReturnsSecondArray_WhenFirstIsEmpty()
    {
        // Arrange
        int[] array01 = { };
        int[] array02 = { 4, 5, 6 };
        int[] expected = { 4, 5, 6 };

        // Act
        int[] result = MergeOrdenado.Cli.Program.MergeOrdenado(array01, array02);

        // Assert
        Assert.Equal(expected, result);
    }

    [Fact]
    public void MergeOrdenado_ReturnsEmpty_WhenBothArraysAreEmpty()
    {
        // Arrange
        int[] array01 = { };
        int[] array02 = { };
        int[] expected = { };

        // Act
        int[] result = MergeOrdenado.Cli.Program.MergeOrdenado(array01, array02);

        // Assert
        Assert.Equal(expected, result);
    }

    [Fact]
    public void MergeOrdenado_MergesArraysWithDuplicates()
    {
        // Arrange
        int[] array01 = { 1, 2, 2, 3 };
        int[] array02 = { 2, 2, 4 };
        int[] expected = { 1, 2, 2, 2, 2, 3, 4 };

        // Act
        int[] result = MergeOrdenado.Cli.Program.MergeOrdenado(array01, array02);

        // Assert
        Assert.Equal(expected, result);
    }
}
