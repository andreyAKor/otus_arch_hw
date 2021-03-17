<?php

namespace classes\matrix;

/**
 * Class MatrixGenarator
 * @package classes\matrix
 */
final class MatrixAddition extends AbstractTemplate
{
    /**
     * @param array $matrix
     * @return array
     */
    public function Do(array $matrix): array
    {
        foreach ($matrix as $i => $row) {
            foreach ($row as $j => $elem) {
                $matrix[$i][$j] = $elem + $elem;
            }
        }

        return $matrix;
    }
}
