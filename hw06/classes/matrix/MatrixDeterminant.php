<?php

namespace classes\matrix;

/**
 * Class MatrixDeterminant
 * @package classes\matrix
 */
final class MatrixDeterminant extends AbstractTemplate
{
    /**
     * @param array $matrix
     * @return array
     */
    public function Do(array $matrix): array
    {
        if (!$this->isSquare($matrix)) {
            throw new Exception('Matrix is not square matrix');
        }

        return [$this->getDeterminant($matrix)];
    }

    /**
     * @param array $matrix
     * @return bool
     */
    private function isSquare(array $matrix): bool
    {
        for ($i = 0; $i < count($matrix); ++$i) {
            if (count($matrix[$i]) !== count($matrix)) {
                return false;
            }
        }

        return true;
    }

    /**
     * @param array $matrix
     * @return float
     */
    private function getDeterminant(array $matrix): float
    {
        $n = count($matrix) - 1;
        $determinant = 0;

        // Если минор дошел до размера 1*1
        if ($n == 0) {
            $determinant += $matrix[0][0];
        }

        // Если минор дошел до размера 2*2
        if ($n == 1) {
            $determinant = $determinant + ($matrix[0][0] * $matrix[1][1]) - ($matrix[0][1] * $matrix[1][0]);
        }

        // Если минор дошел до размера 3*3 и более
        if ($n > 1) {
            for ($a = 0; $a < 1; $a++) {
                for ($b = 0; $b < $n || $b == $n; $b++) {
                    $minor = $this->minor($matrix, $n, $a, $b);
                    $determinant = $determinant + $matrix[0][$b] * pow((-1), 1 + $b + 1) * $this->getDeterminant($minor);
                }
            }
        }

        return $determinant;
    }

    /**
     * @param array $matrix
     * @param int $n
     * @param int $i
     * @param int $j
     * @return array
     */
    private function minor(array $matrix, int $n, int $i, int $j): array
    {
        $r = [];

        for ($a = 0; $a < $n || $a == $n; $a++) {
            for ($b = 0; $b < $n || $b == $n; $b++) {
                if ($a < $i && $b < $j) {
                    $r[$a][$b] = $matrix[$a][$b];
                }

                if ($a < $i && $b > $j) {
                    $r[$a][$b - 1] = $matrix[$a][$b];
                }

                if ($a > $i && $b < $j) {
                    $r[$a - 1][$b] = $matrix[$a][$b];
                }

                if ($a > $i && $b > $j) {
                    $r[$a - 1][$b - 1] = $matrix[$a][$b];
                }
            }
        }

        return $r;
    }
}
