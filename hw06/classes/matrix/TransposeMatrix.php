<?php

namespace classes\matrix;

/**
 * Class TransposeMatrix
 * @package classes\matrix
 */
final class TransposeMatrix extends AbstractTemplate
{
    /**
     * @param array $matrix
     * @return array
     */
    public function Do(array $matrix): array
    {
        foreach ($matrix as $i => $row) {
            foreach ($row as $j => $elem) {
                $matrix[$j][$i] = $elem;
            }
        }

        return $matrix;
    }
}
