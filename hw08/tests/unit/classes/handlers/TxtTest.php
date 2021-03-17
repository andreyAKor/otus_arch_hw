<?php namespace classes\handlers;

use \Codeception\Test\Unit;
use classes\reader\ReaderInterface;
use classes\writer\WriterInterface;
use Mockery;

/**
 * Class TxtTest
 * @package classes\handlers
 */
class TxtTest extends Unit
{
    private const READ_FILE = 'file.txt';
    private const READ_CONTENT = 'some txt content';
    private const OUT_FILE = 'out_file.txt';

    /**
     *
     */
    public function testDo()
    {
        $reader = Mockery::mock(ReaderInterface::class);
        $reader->allows()
            ->Read(self::READ_FILE)
            ->andReturns([self::READ_CONTENT]);

        $writer = Mockery::mock(WriterInterface::class);
        $writer->allows()
            ->Write(self::OUT_FILE, [self::READ_CONTENT]);

        $handler = new Txt($reader, $writer, self::OUT_FILE);
        $handler->handle(self::READ_FILE);
    }
}