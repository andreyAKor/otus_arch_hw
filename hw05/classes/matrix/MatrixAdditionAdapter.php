<?php

namespace classes\matrix;

/**
 * Class MatrixAdditionAdapter
 * @package classes\matrix
 */
final class MatrixAdditionAdapter implements MatrixOperationInterface
{
    /**
     *
     */
    private const PROG1 = './prog1.php';

    /**
     *
     */
    private const PHP = 'php';

    /**
     * @var string
     */
    private string $inFile;

    /**
     * @var string
     */
    private string $outFile;

    /**
     * MatrixAdditionAdapter constructor.
     * @param string $inFile
     * @param string $outFile
     */
    public function __construct(string $inFile, string $outFile)
    {
        $this->inFile = $inFile;
        $this->outFile = $outFile;
    }

    /**
     *
     */
    public function Run(): void
    {
        $exec = sprintf('%s %s %s %s', self::PHP, self::PROG1, $this->inFile, $this->outFile);
        $output = [];
        $resultCode = 0;
        exec($exec, $output, $resultCode);

        if ($resultCode) {
            throw new \Exception(sprintf('Programm execution fail "%s": %s', $exec, $resultCode));
        }
    }
}
